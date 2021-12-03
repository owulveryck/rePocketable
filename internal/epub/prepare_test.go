package epub

import (
	"bytes"
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

func TestPreProcess_figure(t *testing.T) {
	var sample = `<figure class="ja jb jc jd je jf cw cx paragraph-image"> <div role="button" tabindex="0" class="jg jh ji jj aj jk"> <div class="cw cx iz"> <div class="jq s ji jr"> <div class="js jt s"> <div class="jl jm t u v jn aj at jo jp"> <img alt="" class="t u v jn aj ju jv jw" src="https://miro.medium.com/max/60/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg?q=20" width="700" height="590" role="presentation" /> </div> <img alt="" class="jl jm t u v jn aj c" width="700" height="590" role="presentation" /><noscript><img alt="" class="t u v jn aj" src="https://miro.medium.com/max/1400/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg" width="700" height="590" srcSet="https://miro.medium.com/max/552/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 276w, https://miro.medium.com/max/1104/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 552w, https://miro.medium.com/max/1280/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 640w, https://miro.medium.com/max/1400/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 700w" sizes="700px" role="presentation" /></noscript> </div> </div> </div> </div> </figure>`

	content := bytes.NewBufferString(sample)
	n, err := html.Parse(content)
	if err != nil {
		t.Fatal(err)
	}
	err = preProcess(n)
	if err != nil {
		t.Fatal(err)
	}
	var b bytes.Buffer
	html.Render(&b, n)
	expected := `<html><head></head><body><figure class="ja jb jc jd je jf cw cx paragraph-image"> <div role="button" tabindex="0" class="jg jh ji jj aj jk"> <div class="cw cx iz"> <div class="jq s ji jr"> <div class="js jt s"> <div class="jl jm t u v jn aj at jo jp">  </div> <img alt="" class="t u v jn aj" src="https://miro.medium.com/max/1400/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg" width="700" height="590" srcset="https://miro.medium.com/max/552/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 276w, https://miro.medium.com/max/1104/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 552w, https://miro.medium.com/max/1280/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 640w, https://miro.medium.com/max/1400/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 700w" sizes="700px" role="presentation"/><noscript><img alt="" class="t u v jn aj" src="https://miro.medium.com/max/1400/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg" width="700" height="590" srcSet="https://miro.medium.com/max/552/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 276w, https://miro.medium.com/max/1104/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 552w, https://miro.medium.com/max/1280/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 640w, https://miro.medium.com/max/1400/1*RSH2vh_xgQtjB68Zb7oBaA.jpeg 700w" sizes="700px" role="presentation" /></noscript> </div> </div> </div> </div> </figure></body></html>`
	if b.String() != expected {
		t.Fail()
	}
}
