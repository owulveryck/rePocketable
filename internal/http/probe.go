package http

import (
	"context"
	"net/http"
	"time"
)

// WaitOnline check if the service is online, every tick; it returns when the service is up
// StartProbe only returns when the provider is reachable or in case of context cancelation
func (d *Downloader) WaitOnline(ctx context.Context, u string) error {
	if d.isOnline(ctx, u) {
		return nil
	}
	liveness := time.NewTicker(d.LivenessCheck)
	timeout := time.NewTicker(d.ProbeTimeout)
	var cancel context.CancelFunc
	ctx, cancel = context.WithCancel(ctx)
	defer cancel()
	defer liveness.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-liveness.C:
			if d.isOnline(ctx, u) {
				return nil
			}
		case <-timeout.C:
			cancel()
		}
	}
}

func (d *Downloader) isOnline(ctx context.Context, u string) bool {
	headRequest, err := http.NewRequestWithContext(ctx, http.MethodHead, u, nil)
	if err != nil {
		return false
	}
	_, err = d.HTTPClient.Do(headRequest)
	return err == nil
}
