package http

import (
	"net"
	"net/http"

	"github.com/kelseyhightower/envconfig"
)

func (d *Downloader) Usage() {
	envconfig.Usage("DOWNLOADER", d)
}

// Configure the provider with environment variables
func NewDownloader() (*Downloader, error) {
	d := &Downloader{}
	err := envconfig.Process("DOWNLOADER", d)
	if err != nil {
		envconfig.Usage("DOWNLOADER", d)
		return nil, err
	}

	// Create the default client
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: d.TransportTimeout,
		}).Dial,
		TLSHandshakeTimeout: d.TransportTimeout,
	}

	d.HTTPClient = &http.Client{
		Timeout:   d.HTTPTimeout,
		Transport: netTransport,
	}
	return d, nil
}
