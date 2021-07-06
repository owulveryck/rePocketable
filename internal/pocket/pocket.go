package pocket

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/motemen/go-pocket/auth"
	"github.com/owulveryck/rePocketable/internal/docs"
	"github.com/owulveryck/rePocketable/internal/http"
)

const (
	prefix = "POCKET"
)

func (p *Pocket) Doc(w io.Writer) {
	docs.Usage(prefix, p, w)

}
func (p *Pocket) Usage() {
	envconfig.Usage(prefix, p)

}

type Pocket struct {
	UpdateFrequency time.Duration `envconfig:"UPDATE_FREQUENCY" default:"1h" required:"true" desc:"How often to query getPocket"`
	HealthCheck     time.Duration `envconfig:"HEALTH_CHECK" default:"30s" required:"true"`
	URL             string        `envconfig:"POCKET_URL" default:"https://getpocket.com/v3/get" required:"true"`
	ConsumerKey     string        `envconfig:"CONSUMER_KEY" required:"true" desc:"See https://getpocket.com/developer/apps/ to get a consumer key"`
	Username        string        `envconfig:"USERNAME" desc:"The pocket username (will try to fetch it if not found)"`
	Token           string        `envconfig:"TOKEN" desc:"The access token, will try to fetch it if not found or invalid"`
	downloader      *http.Downloader
	ItemsC          chan Item `ignored:"true"`
	auth            *auth.Authorization
}

func NewPocket(downloader *http.Downloader) (*Pocket, error) {
	p := &Pocket{}
	err := envconfig.Process(prefix, p)
	if err != nil {
		envconfig.Usage(prefix, p)
		return nil, err
	}
	p.downloader = downloader
	p.ItemsC = make(chan Item)
	err = p.getAccessToken()
	if err != nil {
		return nil, err
	}
	return p, nil
}

func hasSlept(ctx context.Context, healthCheck time.Duration, sleptTime time.Duration) <-chan struct{} {
	signalC := make(chan struct{})
	go func(chan<- struct{}) {
		tick := time.NewTicker(healthCheck)
		last := time.Now()
		for {
			select {
			case <-ctx.Done():
				return
			case <-tick.C:
				if time.Now().Round(0).Sub(last) > sleptTime {
					signalC <- struct{}{}
				}
				last = time.Now()
			}
		}
	}(signalC)
	return signalC
}

func (p *Pocket) RunPoller(ctx context.Context) error {
	ticker := time.NewTicker(p.UpdateFrequency)
	err := p.retrieveArticles(ctx, ticker)
	if err != nil {
		log.Fatal(err)
	}
	runC := make(chan struct{}, 1)
	sleptC := hasSlept(ctx, p.HealthCheck, p.HealthCheck*2)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			runC <- struct{}{}
		case <-sleptC:
			runC <- struct{}{}
			ticker.Reset(p.UpdateFrequency)
		case <-runC:
			err := p.retrieveArticles(ctx, ticker)
			if err != nil {
				return err
			}
		}
	}
}

func (p *Pocket) retrieveArticles(ctx context.Context, ticker *time.Ticker) error {
	err := p.downloader.WaitOnline(ctx, p.URL)
	if err != nil {
		return err
	}
	ticker.Reset(p.UpdateFrequency)
	err = p.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
