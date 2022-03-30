package epub

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bmaupin/go-epub"
	"github.com/cixtor/readability"
	"github.com/dyatlov/go-opengraph/opengraph"
	"golang.org/x/net/html"
)

func init() {
	epub.Use(epub.MemoryFS)
}

type Document struct {
	*epub.Epub
	Element            string
	buf                strings.Builder
	CurrentSectionName string
	Client             *http.Client
	CSS                io.Reader
	OG                 *opengraph.OpenGraph
}

func NewDocument(element string) *Document {
	return &Document{
		Epub:               epub.NewEpub(""),
		Element:            element,
		buf:                strings.Builder{},
		CurrentSectionName: "Preamble",
	}
}

func (d *Document) Fill(ctx context.Context) error {
	client := http.DefaultClient
	if d.Client != nil {
		d.Epub.Client = d.Client
		client = d.Client
	}
	r := readability.New()
	req, err := http.NewRequestWithContext(ctx, "GET", d.Element, nil)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("cannot fetch document: %w", err)
	}
	defer res.Body.Close()

	og, content := getOpenGraph(res.Body)
	firstPass := false
	if d.OG == nil {
		d.OG = og
		firstPass = true
	}
	doc, err := html.Parse(content)
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
	article, err := r.Parse(pipeR, d.Element)
	if err != nil {
		return fmt.Errorf("cannot parse document: %w", err)
	}
	if firstPass {
		err = d.setMeta(&article)
		if err != nil {
			return err
		}
	}
	err = d.replaceImages(article.Node)
	if err != nil {
		return err
	}
	css, err := d.setCSS()
	if err != nil {
		log.Println(err)
	}
	if firstPass {
		d.createMeta()
	}

	var body strings.Builder
	err = html.Render(&body, article.Node)
	if err != nil {
		return err
	}

	_, err = d.AddSection(body.String(), d.CurrentSectionName, "", css)
	return err
}

func (d *Document) setMeta(a *readability.Article) error {
	d.SetTitle(a.Title)
	d.SetDescription(d.OG.Description)
	d.SetAuthor(a.Byline)
	if a.Image != "" {
		img, err := imageToCover(a.Image, a.Title, a.Byline, a.SiteName)
		if err != nil {
			return err
		}
		img, err = d.AddImage(img, "")
		if err != nil {
			return err
		}
		d.SetCover(img, "")
	}
	return nil
}

func (d *Document) getURL(attr []html.Attribute) (source string, filename string, err error) {
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
		if a.Key == "data-src" {
			u, err := url.Parse(a.Val)
			if err != nil {
				return "", "", err
			}
			scheme = u.Scheme
			host = u.Host
			if val == "" {
				val = a.Val
			}
			attr[i].Key = "src"
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
	// if no scheme, force https
	if u.Scheme == "" {
		ru, _ := url.Parse(d.Element)
		u.Scheme = ru.Scheme

	}
	if u.Host == "" {
		ru, _ := url.Parse(d.Element)
		u.Host = ru.Host
	}
	f := fmt.Sprintf("%x.%v", md5.Sum([]byte(u.String())), filepath.Ext(val))
	return u.String(), f, nil
}

func (d *Document) replaceImages(n *html.Node) error {
	if n.Type == html.ElementNode && n.Data == "img" {
		val, f, err := d.getURL(n.Attr)
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
