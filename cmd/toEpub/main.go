package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/owulveryck/rePocketable/internal/epub"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

var DB map[string]pocket.Item

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	downloader, err := http.NewDownloader()
	if err != nil {
		log.Fatal(err)
	}

	item := pocket.Item{
		ResolvedURL: os.Args[1],
		GivenURL:    os.Args[1],
	}
	doc := epub.NewDocument(item)
	doc.Client = downloader.HTTPClient
	doc.CSS = css
	err = doc.Fill(ctx)
	if err != nil {
		log.Println("Cannot fill document: ", err)
		return
	}
	log.Println("writing output")
	err = doc.Write(fmt.Sprintf("%v.epub", filepath.Base(os.Args[1])))
	if err != nil {
		log.Fatal("Cannot write document: ", err)
	}
}
