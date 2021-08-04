package epub

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"path/filepath"
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
	article, err := r.Parse(res.Body, d.item.ResolvedURL)
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

func (d *Document) replaceImages(n *html.Node) error {
	if n.Type == html.ElementNode && n.Data == "img" {
		for i, a := range n.Attr {
			if a.Key == "src" {
				// get the filname
				u, err := url.Parse(a.Val)
				if err != nil {
					return err
				}
				f := filepath.Base(u.Path)
				img, err := d.AddImage(a.Val, f)
				// if err, try to download it again with default name
				if err != nil {
					img, err = d.AddImage(a.Val, "")
					if err != nil {
						return err
					}
				}
				n.Attr[i].Val = img
			}
			// remove the srcset
			if a.Key == "srcset" {
				n.Attr[i] = n.Attr[len(n.Attr)-1]        // Copy last element to index i.
				n.Attr[len(n.Attr)-1] = html.Attribute{} // Erase last element (write zero value).
				n.Attr = n.Attr[:len(n.Attr)-1]          // Truncate slice.
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
