package epub

import (
	"bytes"
	"io"
	"log"
	"regexp"

	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/vincent-petithory/dataurl"
	"golang.org/x/net/html"
)

const (
	size = float64(12)
	dpi  = float64(72)
)

var mathJax = regexp.MustCompile(`\$\$[^\$]+\$\$`)

func preProcess(n *html.Node) error {
	switch {
	case n.Type == html.ElementNode && n.Data == "figure":
		processFigure(n)
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
	log.Println("Processing", n.Data)
	log.Println(mathJax.FindAllString(n.Data, -1))
	return nil
	fnts := lmromanFonts()
	var b bytes.Buffer
	dst := drawimg.NewRenderer(&b)
	err := mtex.Render(dst, n.FirstChild.Data, size, dpi, fnts)
	if err != nil {
		return err
	}
	dataURL := dataurl.New(b.Bytes(), "image/png")
	content, err := dataURL.MarshalText()
	if err != nil {
		return err
	}
	n.Data = "img"
	n.Attr = []html.Attribute{
		{
			Key: "src",
			Val: string(content),
		},
	}

	return nil
}

func hasAttribute(n *html.Node, attr html.Attribute) bool {
	log.Println(n.Attr)
	for i := 0; i < len(n.Attr); i++ {
		if n.Attr[i].Key == attr.Key && n.Attr[i].Val == attr.Val {
			return true
		}
	}
	return false
}

// math/tex; mode=display

func processFigure(n *html.Node) error {
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
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := processFigure(c)
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
