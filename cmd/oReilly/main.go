package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	nethttp "net/http"
	"net/url"

	"github.com/owulveryck/rePocketable/internal/epub"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
	"github.com/robertkrimen/otto/ast"
	"github.com/robertkrimen/otto/file"
	"github.com/robertkrimen/otto/parser"
	"golang.org/x/net/html"
)

var DB map[string]pocket.Item

type headers map[string][]string

func (h headers) String() string {
	var b strings.Builder
	for k, v := range h {
		fmt.Fprintf(&b, "%v: %v|", k, v)
	}
	return b.String()
}

func (h headers) Set(v string) error {
	elements := strings.SplitN(v, ":", 2)
	if len(elements) != 2 {
		return errors.New("bad header passed")
	}
	h[elements[0]] = append(h[elements[0]], elements[1])
	return nil
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var headersFlag headers
	headersFlag = make(map[string][]string)
	flag.Var(&headersFlag, "H", "header")
	flag.Parse()

	downloader, err := http.NewDownloader(nethttp.Header(headersFlag))
	if err != nil {
		log.Fatal(err)
	}
	content := os.Args[len(os.Args)-1]
	u, err := url.Parse(content)
	if err != nil {
		log.Fatal(err)
	}
	res, err := downloader.Get(ctx, content)
	if err != nil {
		log.Fatal(err)
	}
	d, err := html.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	t := &tocRetriever{}
	err = t.preProcess(d)
	if err != nil {
		log.Fatal(err)
	}

	doc := epub.NewDocument("")
	doc.Client = downloader.HTTPClient
	doc.CSS = css
	for _, url := range t.urls {
		u2 := *u
		u2.Path = url
		log.Println(u2.String())
		doc.Element = u2.String()
		doc.CurrentSectionName = u2.String()
		err = doc.Fill(ctx)
		if err != nil {
			log.Println("Cannot fill document: ", err)
			return
		}
	}
	log.Println("writing output: ", fmt.Sprintf("%v.epub", filepath.Base(os.Args[len(os.Args)-1])))
	err = doc.Write(fmt.Sprintf("%v.epub", filepath.Base(os.Args[len(os.Args)-1])))
	if err != nil {
		log.Fatal("Cannot write document: ", err)
	}
}

type tocRetriever struct {
	urls []string
}

func (tr *tocRetriever) preProcess(n *html.Node) error {
	switch {
	case n.Type == html.ElementNode && n.Data == "script":
		if n.FirstChild != nil {
			content := n.FirstChild.Data
			program, err := parser.ParseFile(nil, "", content, 0)
			if err != nil {
				return err
			}
			w := &walkExample{}

			ast.Walk(w, program)
			if w.left != 0 && w.right != 0 {
				var t Toc
				err := json.Unmarshal([]byte(content[w.left-1:w.right]), &t)
				if err != nil {
					return err
				}
				for _, v := range t.Appstate.Tableofcontents {
					for _, s := range v.Sections {
						if tr.urls == nil {
							tr.urls = make([]string, 1)
							tr.urls[0] = s.Apiurl
							continue
						}
						if tr.urls[len(tr.urls)-1] == s.Apiurl {
							continue
						}
						tr.urls = append(tr.urls, s.Apiurl)
					}
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := tr.preProcess(c)
		if err != nil {
			return err
		}
	}
	return nil
}

type walkExample struct {
	shift       file.Idx
	right, left int
}

func (w *walkExample) Enter(n ast.Node) ast.Visitor {
	if id, ok := n.(*ast.AssignExpression); ok && id != nil {
		if key, ok := id.Left.(*ast.DotExpression); ok && key != nil {
			if key.Identifier.Name == "initialStoreData" {
				w.left = int(id.Right.Idx0())
				w.right = int(id.Right.Idx1())
			}
		}
	}

	return w
}

func (w *walkExample) Exit(n ast.Node) {
	// AST node n has had all its children walked. Pop it out of your
	// stack, or do whatever processing you need to do, if any.
}
