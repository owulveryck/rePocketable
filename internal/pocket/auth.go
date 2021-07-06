package pocket

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/motemen/go-pocket/auth"
)

// getAccessToken from the cache, try to fetch one if not present
func (p *Pocket) getAccessToken() error {
	p.auth = &auth.Authorization{
		AccessToken: p.Token,
		Username:    p.Username,
	}
	if p.Token == "" || p.Username == "" {
		accessToken, err := obtainAccessToken(p.ConsumerKey)
		if err != nil {
			return err
		}
		p.auth = accessToken
		log.Printf("POCKET_TOKEN=%v POCKET_USERNAME=%v", p.auth.AccessToken, p.auth.Username)
	}
	return nil
}

func obtainAccessToken(consumerKey string) (*auth.Authorization, error) {
	ch := make(chan struct{})
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			if req.URL.Path == "/favicon.ico" {
				http.Error(w, "Not Found", 404)
				return
			}

			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintln(w, "Authorized.")
			ch <- struct{}{}
		}))
	defer ts.Close()

	redirectURL := ts.URL

	requestToken, err := auth.ObtainRequestToken(consumerKey, redirectURL)
	if err != nil {
		return nil, err
	}

	url := auth.GenerateAuthorizationURL(requestToken, redirectURL)
	fmt.Println(url)

	<-ch

	return auth.ObtainAccessToken(consumerKey, requestToken)
}
