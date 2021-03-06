package epub

import (
	"bytes"
	"io"
	"regexp"

	"golang.org/x/net/html"
)

const (
	size = float64(14)
	dpi  = float64(96)
)

var mathJax = regexp.MustCompile(`\$\$?[^\$]+\$?\$`)

func preProcess(n *html.Node) error {
	switch {
	case n.Type == html.ElementNode && n.Data == "figure":
		f := &figure{
			images: make([]*html.Node, 0),
		}
		f.processFigure(n)
		// Clear all other images (medium, towarddatascience, ...)
		if len(f.images) > 1 {
			for _, img := range f.images {
				if img != f.validImage {
					img.Parent.RemoveChild(img)
				}
			}
		}
	case n.Type == html.TextNode && hasMathJax(n):
		processMathTex(n, false)
		processMathTex(n, true)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if n.Type == html.CommentNode || (n.Type == html.ElementNode && n.Data == "script") {
			continue
		}
		err := preProcess(c)
		if err != nil {
			return err
		}
	}
	return nil
}

type figure struct {
	images     []*html.Node
	validImage *html.Node
}

func (f *figure) processFigure(n *html.Node) error {
	if n.Type == html.ElementNode && n.Data == "img" {
		f.images = append(f.images, n)
	}
	if n.Data == "noscript" {
		if originalImg := n.PrevSibling; originalImg.Data == "img" {
			// the img data is encoded as a string in the n.FirstChild.Data field
			// Let's parse it as a node:
			doc, err := html.Parse(bytes.NewBufferString(n.FirstChild.Data))
			if err != nil {
				return err
			}
			img := getImgNode(doc)
			if img != nil {
				originalImg.Attr = img.Attr
			}
			f.validImage = originalImg
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := f.processFigure(c)
		if err != io.EOF {
			return err
		}
	}
	return io.EOF
}

func getImgNode(node *html.Node) *html.Node {
	if node.Type == html.ElementNode && node.Data == "img" {
		return node
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		n := getImgNode(child)
		if n != nil {
			return n
		}
	}
	return nil
}
