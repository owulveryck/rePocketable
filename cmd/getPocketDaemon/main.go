package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/owulveryck/rePocketable/internal/epub"
	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

var DB map[string]pocket.Item

func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	dstDir := flag.String("dst", "testdata", "destination directory")
	dumpDB := flag.Bool("dump", false, "dump the content of the DB on stdout")
	dumpDBHTML := flag.Bool("dumphtml", false, "dump the content of the DB on stdout")
	if usage() {
		return
	}
	if *dumpDB {
		err := dumpContent(os.Stdout, filepath.Join(*dstDir, ".db"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if *dumpDBHTML {
		err := dumpHTMLContent(os.Stdout, filepath.Join(*dstDir, ".db"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	database := initDatabase(sigs, cancel, *dstDir)
	downloader, err := http.NewDownloader(nil)
	if err != nil {
		log.Fatal(err)
	}
	myPocket, err := pocket.NewPocket(downloader)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err = myPocket.RunPoller(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()
	maxNbConcurrentGoroutines := 15
	concurrentGoroutines := make(chan struct{}, maxNbConcurrentGoroutines)
	// Fill the dummy channel with maxNbConcurrentGoroutines empty struct.
	for i := 0; i < maxNbConcurrentGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}
	for item := range myPocket.ItemsC {
		if item.IsArticle != 0 {
			<-concurrentGoroutines
			go func(item pocket.Item) {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Recovered:", r)
					}
					database.Store(item.ItemID, item)
					// Say that another goroutine can now start.
					concurrentGoroutines <- struct{}{}
				}()
				log.Println("processing ", item.ItemID)
				_, ok := database.Load(item.ItemID)
				if ok {
					log.Printf("%v already present (%v)", item.ItemID, item.ResolvedTitle)
					return
				}
				doc := epub.NewDocument(item)
				doc.Client = downloader.HTTPClient
				err := doc.Fill(ctx)
				if err != nil {
					log.Println("Cannot fill document: ", err)
					return
				}
				err = doc.Write(filepath.Join(*dstDir, fmt.Sprintf("%v.epub", item.ItemID)))
				if err != nil {
					log.Println("Cannot write document: ", err)
					return
				}
			}(item)
		}
	}
}
