package main

import (
	"flag"
	"os"

	"github.com/owulveryck/rePocketable/internal/http"
	"github.com/owulveryck/rePocketable/internal/pocket"
)

func usage() bool {
	help := flag.Bool("h", false, "help")
	doc := flag.Bool("d", false, "geretate usage for documentation (MD)")
	flag.Parse()
	if *help {
		d := &http.Downloader{}
		d.Usage()
		p := &pocket.Pocket{}
		p.Usage()
		return true
	}
	if *doc {
		d := &http.Downloader{}
		d.Doc(os.Stdout)
		p := &pocket.Pocket{}
		p.Doc(os.Stdout)
		return true
	}
	return false
}
