//go:build js && wasm
// +build js,wasm

package http

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/owulveryck/rePocketable/internal/docs"
)

const (
	prefix = "DOWNLOADER"
)

func (d *Downloader) Doc(w io.Writer) {
	docs.Usage(prefix, d, w)

}

func (d *Downloader) Usage() {
	envconfig.Usage(prefix, d)
}

// Configure the provider with environment variables
func NewDownloader() (*Downloader, error) {
	d := &Downloader{}
	err := envconfig.Process(prefix, d)
	if err != nil {
		envconfig.Usage(prefix, d)
		return nil, err
	}

	// Create the default client
	var netTransport = &corsCompatibleTransport{
		&http.Transport{
			Dial: (&net.Dialer{
				Timeout: d.TransportTimeout,
			}).Dial,
			TLSHandshakeTimeout: d.TransportTimeout,
		}}

	d.HTTPClient = &http.Client{
		Timeout:   d.HTTPTimeout,
		Transport: netTransport,
	}
	return d, nil
}

type corsCompatibleTransport struct {
	*http.Transport
}

func (c *corsCompatibleTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Println("setting fetch mode to no-cors")
	req.Header.Set("js.fetch:mode", "no-cors")
	return c.Transport.RoundTrip(req)
}
