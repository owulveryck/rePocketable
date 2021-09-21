package epub

import (
	"bytes"
	"io"
	"log"
	"sync"

	"github.com/dyatlov/go-opengraph/opengraph"
)

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
