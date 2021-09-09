package epub

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmaupin/go-epub"
	"github.com/cixtor/readability"
	"github.com/owulveryck/rePocketable/internal/pocket"
	"golang.org/x/net/html"
)

type Document struct {
	*epub.Epub
	item        pocket.Item
	buf         strings.Builder
	currSection string
	Client      *http.Client
}

func NewDocument(item pocket.Item) *Document {
	return &Document{
		Epub:        epub.NewEpub(""),
		item:        item,
		buf:         strings.Builder{},
		currSection: "Preamble",
	}
}

func (d *Document) Fill(ctx context.Context) error {
	client := http.DefaultClient
	if d.Client != nil {
		client = d.Client
	}
	r := readability.New()
	req, err := http.NewRequestWithContext(ctx, "GET", d.item.URL(), nil)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("cannot fetch document: %w", err)
	}
	defer res.Body.Close()
	doc, err := html.Parse(res.Body)
	if err != nil {
		return err
	}
	err = preProcess(doc)
	if err != nil {
		log.Fatal(err)
	}
	//article, err := r.Parse(res.Body, d.item.ResolvedURL)
	pipeR, pipeW := io.Pipe()
	go func() {
		defer pipeW.Close()
		err = html.Render(pipeW, doc)
		if err != nil {
			return
		}
	}()
	html.Render(os.Stdout, doc)
	article, err := r.Parse(pipeR, d.item.ResolvedURL)
	if err != nil {
		return fmt.Errorf("cannot parse document: %w", err)
	}
	err = d.setMeta(&article)
	if err != nil {
		return err
	}
	err = d.replaceImages(article.Node)
	if err != nil {
		return err
	}
	var body strings.Builder
	err = html.Render(&body, article.Node)
	if err != nil {
		return err
	}
	_, err = d.AddSection(body.String(), "Content", "", "")
	return err
}

func (d *Document) setMeta(a *readability.Article) error {
	d.SetTitle(d.item.ResolvedTitle)
	d.SetDescription(d.item.Excerpt)
	d.SetAuthor(a.Byline)
	if a.Image != "" {
		img, err := d.AddImage(a.Image, "")
		if err != nil {
			return err
		}
		d.SetCover(img, "")
	}
	return nil
}

func getURL(attr []html.Attribute) (source string, filename string, err error) {
	var val string
	var host, scheme string
	for i := 0; i < len(attr); i++ {
		a := attr[i]
		if a.Key == "src" {
			u, err := url.Parse(a.Val)
			if err != nil {
				return "", "", err
			}
			scheme = u.Scheme
			host = u.Host
			if val == "" {
				val = a.Val
			}
		}
		if a.Key == "srcset" {
			set, err := csv.NewReader(bytes.NewBufferString(a.Val)).Read()
			if err != nil {
				return "", "", err
			}
			srcSet, err := newSrcSetElementsFromStrings(set)
			if err != nil {
				return "", "", err
			}
			sort.Sort(srcSet)
			val = srcSet[0].url
		}
	}
	// get the filname
	u, err := url.Parse(val)
	if err != nil {
		return "", "", err
	}
	// if no scheme, force https
	if u.Scheme == "" {
		u.Scheme = scheme
	}
	if u.Host == "" {
		u.Host = host
	}
	f := filepath.Base(u.Path)
	return u.String(), f, nil
}

func (d *Document) replaceImages(n *html.Node) error {
	if n.Type == html.ElementNode && n.Data == "img" {
		val, f, err := getURL(n.Attr)
		if err != nil {
			return err
		}
		for i, a := range n.Attr {
			if a.Key == "src" {
				img, err := d.AddImage(val, f)
				// if err, try to download it again with default name
				if err != nil {
					img, err = d.AddImage(val, "")
					if err != nil {
						return err
					}
				}
				n.Attr[i].Val = img
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := d.replaceImages(c)
		if err != nil {
			return err
		}
	}
	return nil
}
