package http

import (
	"net/http"
	"time"
)

// Downloader
type Downloader struct {
	LivenessCheck    time.Duration `envconfig:"LIVENESS_CHECK" default:"5m" required:"true"`
	ProbeTimeout     time.Duration `envconfig:"PROBE_TIMEOUT" default:"60m" required:"true"`
	HTTPTimeout      time.Duration `envconfig:"HTTP_TIMEOUT" default:"10s" required:"true"`
	TransportTimeout time.Duration `envconfig:"TRANSPORT_TIMEOUT" default:"5s" required:"true"`
	HTTPClient       *http.Client  `ignored:"true"`
}
