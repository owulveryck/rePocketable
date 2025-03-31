package markdown

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/cixtor/readability"
	"github.com/owulveryck/rePocketable/internal/pocket"
	"golang.org/x/net/html"
	"github.com/dyatlov/go-opengraph/opengraph"
)

// Document represents a markdown document with content extracted from a URL
type Document struct {
	Title       string
	Description string
	Author      string
	Content     string
	URL         string
	Client      *http.Client
	OG          *opengraph.OpenGraph
}

// NewDocument creates a new markdown document from a pocket item
func NewDocument(item pocket.Item) *Document {
	return &Document{
		URL: item.URL(),
	}
}

// Fill populates the document with content from the URL
func (d *Document) Fill(ctx context.Context) error {
	client := http.DefaultClient
	if d.Client != nil {
		client = d.Client
	}
	
	r := readability.New()
	req, err := http.NewRequestWithContext(ctx, "GET", d.URL, nil)
	if err != nil {
		return fmt.Errorf("cannot create request: %w", err)
	}
	
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("cannot fetch document: %w", err)
	}
	defer res.Body.Close()

	// Get OpenGraph data
	og, content := getOpenGraph(res.Body)
	d.OG = og
	
	doc, err := html.Parse(content)
	if err != nil {
		return err
	}
	
	err = preProcess(doc)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the document
	pipeR, pipeW := io.Pipe()
	go func() {
		defer pipeW.Close()
		err = html.Render(pipeW, doc)
		if err != nil {
			return
		}
	}()
	
	article, err := r.Parse(pipeR, d.URL)
	if err != nil {
		return fmt.Errorf("cannot parse document: %w", err)
	}
	
	// Set metadata
	d.Title = article.Title
	d.Description = d.OG.Description
	d.Author = article.Byline
	
	// Convert HTML to Markdown
	markdown, err := ConvertHTMLToMarkdown(article.Node)
	if err != nil {
		return fmt.Errorf("cannot convert HTML to markdown: %w", err)
	}
	
	d.Content = markdown
	return nil
}

// Write writes the markdown content to a file
func (d *Document) Write(filename string) error {
	// Create a markdown file with the content
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("cannot create file: %w", err)
	}
	defer f.Close()
	
	return d.WriteTo(f)
}

// WriteTo writes the markdown content to an io.Writer
func (d *Document) WriteTo(w io.Writer) error {
	// Write metadata as frontmatter
	fmt.Fprintf(w, "---\n")
	fmt.Fprintf(w, "title: %s\n", d.Title)
	if d.Author != "" {
		fmt.Fprintf(w, "author: %s\n", d.Author)
	}
	if d.Description != "" {
		fmt.Fprintf(w, "description: %s\n", d.Description)
	}
	fmt.Fprintf(w, "source: %s\n", d.URL)
	fmt.Fprintf(w, "---\n\n")
	
	// Write the content
	fmt.Fprintf(w, "# %s\n\n", d.Title)
	fmt.Fprintf(w, "%s\n", d.Content)
	
	return nil
}

// ConvertHTMLToMarkdown converts an HTML node to markdown text
func ConvertHTMLToMarkdown(n *html.Node) (string, error) {
	var markdown strings.Builder
	
	// Traverse the HTML tree and convert to markdown
	err := traverseHTML(n, &markdown, 0)
	if err != nil {
		return "", err
	}
	
	return markdown.String(), nil
}

// traverseHTML recursively traverses the HTML tree and converts it to markdown
func traverseHTML(n *html.Node, markdown *strings.Builder, depth int) error {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			markdown.WriteString(text)
			if text[len(text)-1] != ' ' {
				markdown.WriteString(" ")
			}
		}
		return nil
	}
	
	if n.Type == html.ElementNode {
		switch n.Data {
		case "h1":
			markdown.WriteString("\n# ")
		case "h2":
			markdown.WriteString("\n## ")
		case "h3":
			markdown.WriteString("\n### ")
		case "h4":
			markdown.WriteString("\n#### ")
		case "h5":
			markdown.WriteString("\n##### ")
		case "h6":
			markdown.WriteString("\n###### ")
		case "p":
			markdown.WriteString("\n\n")
		case "br":
			markdown.WriteString("\n")
		case "strong", "b":
			markdown.WriteString("**")
		case "em", "i":
			markdown.WriteString("*")
		case "a":
			markdown.WriteString("[")
			// We'll close this later and add the URL
		case "ul":
			markdown.WriteString("\n")
		case "ol":
			markdown.WriteString("\n")
		case "li":
			markdown.WriteString("\n- ")
		case "blockquote":
			markdown.WriteString("\n> ")
		case "code":
			markdown.WriteString("`")
		case "pre":
			markdown.WriteString("\n```\n")
		case "img":
			// Skip images for now as requested
			return nil
		}
	}
	
	// Process children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := traverseHTML(c, markdown, depth+1)
		if err != nil {
			return err
		}
	}
	
	// Close tags that need closing
	if n.Type == html.ElementNode {
		switch n.Data {
		case "strong", "b":
			markdown.WriteString("**")
		case "em", "i":
			markdown.WriteString("*")
		case "a":
			markdown.WriteString("](")
			// Find href attribute
			for _, a := range n.Attr {
				if a.Key == "href" {
					markdown.WriteString(a.Val)
					break
				}
			}
			markdown.WriteString(")")
		case "code":
			markdown.WriteString("`")
		case "pre":
			markdown.WriteString("\n```\n")
		}
	}
	
	return nil
}

// getOpenGraph extract the data from the io.Reader and returns a new reader
func getOpenGraph(r io.Reader) (*opengraph.OpenGraph, io.Reader) {
	var buf bytes.Buffer
	og := opengraph.NewOpenGraph()
	pr, pw := io.Pipe()

	// we need to wait for everything to be done
	wg := sync.WaitGroup{}
	wg.Add(2)

	// TeeReader gets the data from the r and also writes it to the PipeWriter
	tr := io.TeeReader(r, pw)

	go func() {
		defer wg.Done()
		defer pw.Close()

		// get data from the TeeReader, which feeds the PipeReader through the PipeWriter
		err := og.ProcessHTML(tr)
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		// read from the PipeReader to stdout
		if _, err := io.Copy(&buf, pr); err != nil {
			log.Fatal(err)
		}
	}()

	wg.Wait()
	return og, &buf
}

// preProcess prepares the HTML for conversion
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
		if originalImg := n.PrevSibling; originalImg != nil && originalImg.Data == "img" {
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