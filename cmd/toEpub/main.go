package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	nethttp "net/http"

	"github.com/owulveryck/rePocketable/internal/epub"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
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

	item := pocket.Item{
		ResolvedURL: os.Args[len(os.Args)-1],
		GivenURL:    os.Args[len(os.Args)-1],
	}
	doc := epub.NewDocument(item)
	doc.Client = downloader.HTTPClient
	doc.CSS = css
	err = doc.Fill(ctx)
	if err != nil {
		log.Println("Cannot fill document: ", err)
		return
	}
	log.Println("writing output: ", fmt.Sprintf("%v.epub", filepath.Base(os.Args[len(os.Args)-1])))
	err = doc.Write(fmt.Sprintf("%v.epub", filepath.Base(os.Args[len(os.Args)-1])))
	if err != nil {
		log.Fatal("Cannot write document: ", err)
	}
}
