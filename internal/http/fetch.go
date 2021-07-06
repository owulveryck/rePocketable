package http

import (
	"context"
	"net/http"
)

func (d *Downloader) Get(ctx context.Context, u string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	return d.HTTPClient.Do(req)
}
