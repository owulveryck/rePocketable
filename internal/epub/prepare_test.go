package epub

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestPreProcess(t *testing.T) {
	f, err := os.Open("testdata/sample.html")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	n, err := html.Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	preProcess(n)

}
