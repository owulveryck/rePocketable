package main

import (
	"context"
	"encoding/gob"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/kr/pretty"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

var DB map[string]pocket.Item

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dstDir := flag.String("dst", "testdata", "destination directory")
	help := flag.Bool("h", false, "help")
	flag.Parse()
	if *help {
		d := &http.Downloader{}
		d.Usage()
		p := &pocket.Pocket{}
		p.Usage()
		return
	}
	f, err := os.Open(filepath.Join(*dstDir, ".db"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	err = dec.Decode(&DB)
	if err != nil {
		log.Println("access db", err)
	}

	downloader, err := http.NewDownloader()
	if err != nil {
		log.Fatal(err)
	}
	pocket, err := pocket.NewPocket(downloader)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		err = pocket.RunPoller(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()
	for item := range pocket.ItemsC {
		if item.IsArticle == 0 {
			pretty.Print(item)
		}
	}
}
