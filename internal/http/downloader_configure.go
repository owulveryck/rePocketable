package http

import (
	"io"
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
func NewDownloader(withHeaders http.Header) (*Downloader, error) {
	d := &Downloader{}
	err := envconfig.Process(prefix, d)
	if err != nil {
		envconfig.Usage(prefix, d)
		return nil, err
	}

	// Create the default client
	var netTransport = &customTransport{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: d.TransportTimeout,
			}).Dial,
			TLSHandshakeTimeout: d.TransportTimeout,
		},
		headers: withHeaders,
	}

	d.HTTPClient = &http.Client{
		Timeout:   d.HTTPTimeout,
		Transport: netTransport,
	}
	return d, nil
}

type customTransport struct {
	headers http.Header
	*http.Transport
}

func (c *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if c.headers != nil {
		req.Header = c.headers
	}
	return c.Transport.RoundTrip(req)
}
