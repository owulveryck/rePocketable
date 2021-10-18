package epub

import (
	"bytes"
	"io"
	"log"
	"regexp"
	"strings"

	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/vincent-petithory/dataurl"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
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
		processMathTex(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := preProcess(c)
		if err != nil {
			return err
		}
	}
	return nil
}

func hasMathJax(n *html.Node) bool {
	return len(mathJax.FindAllString(n.Data, -1)) > 0
}

func processMathTex(n *html.Node) error {
	var currentFormula []byte
	defer func() {
		if r := recover(); r != nil {
			log.Printf("mathjax processing error: recovered %v in %s", r, currentFormula)
		}
	}()
	completeText := n.Data
	//	fnts := lmromanFonts()
	fnts := liberationFonts()

	allFormulas := mathJax.FindAll([]byte(completeText), -1)
	images := make([]*html.Node, len(allFormulas))
	var i int
	var remaining string
	var delete bool
	for i, currentFormula = range allFormulas {
		expr := strings.TrimFunc(string(currentFormula), func(r rune) bool {
			return r == '$'
		})
		var b bytes.Buffer
		dst := drawimg.NewRenderer(&b)
		err := mtex.Render(dst, "$"+expr+"$", size, dpi, fnts)
		if err != nil {
			return err
		}
		dataURL := dataurl.New(b.Bytes(), "image/png")
		content, err := dataURL.MarshalText()
		if err != nil {
			return err
		}
		images[i] = &html.Node{
			Type:      html.ElementNode,
			DataAtom:  atom.Img,
			Data:      "img",
			Namespace: n.Namespace,
			Attr: []html.Attribute{
				{
					Key: "src",
					Val: string(content),
				},
			},
		}
		delete = true
		if remaining == "" {
			remaining = n.Data
		}
		data := strings.SplitN(remaining, string(currentFormula), 2)
		remaining = data[1]
		firstpart := &html.Node{
			Type:      n.Type,
			DataAtom:  n.DataAtom,
			Data:      data[0],
			Namespace: n.Namespace,
			Attr:      n.Attr,
		}
		n.Parent.InsertBefore(images[i], n)
		n.Parent.InsertBefore(firstpart, images[i])
		//n.Parent.AppendChild(images[i])
	}
	if remaining != "" {
		n.Parent.InsertBefore(&html.Node{
			Type:      n.Type,
			DataAtom:  n.DataAtom,
			Data:      remaining,
			Namespace: n.Namespace,
			Attr:      n.Attr,
		}, n)
	}
	if delete {
		n.Data = ""
		//n.Parent.RemoveChild(n)
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
