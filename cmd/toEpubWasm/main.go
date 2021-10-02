package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/owulveryck/rePocketable/internal/epub"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

var DB map[string]pocket.Item

func main() {
	content := "https://blog.owulveryck.info/2021/06/08/pov-a-streaming/communication-platform-for-the-data-mesh.html"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	downloader, err := http.NewDownloader()
	if err != nil {
		log.Fatal(err)
	}

	item := pocket.Item{
		ResolvedURL: content,
		GivenURL:    content,
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
	err = doc.Write(fmt.Sprintf("%v.epub", filepath.Base(content)))
	if err != nil {
		log.Fatal("Cannot write document: ", err)
	}
}
