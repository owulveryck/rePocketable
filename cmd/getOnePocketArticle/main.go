package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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
	myPocket, err := pocket.NewPocket(downloader)
	if err != nil {
		log.Fatal(err)
	}

	maxNbConcurrentGoroutines := 15
	concurrentGoroutines := make(chan struct{}, maxNbConcurrentGoroutines)
	// Fill the dummy channel with maxNbConcurrentGoroutines empty struct.
	for i := 0; i < maxNbConcurrentGoroutines; i++ {
		concurrentGoroutines <- struct{}{}
	}
	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("arg1 must be an int")
	}
	go func() {
		myPocket.Get(ctx)
	}()
	for item := range myPocket.ItemsC {
		if item.ItemID == id {
			<-concurrentGoroutines
			go func(item pocket.Item) {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Recovered:", r)
					}
					concurrentGoroutines <- struct{}{}
					close(myPocket.ItemsC)
				}()
				log.Println("processing ", item.ItemID)
				doc := epub.NewDocument(item)
				doc.Client = downloader.HTTPClient
				doc.CSS = css
				err := doc.Fill(ctx)
				if err != nil {
					log.Println("Cannot fill document: ", err)
					return
				}
				log.Println("writing output: " + fmt.Sprintf("%v.epub", item.ItemID))
				err = doc.Write(fmt.Sprintf("%v.epub", item.ItemID))
				if err != nil {
					log.Println("Cannot write document: ", err)
					return
				}
			}(item)
		}
	}
}
