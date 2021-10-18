package epub

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/owulveryck/rePocketable/internal/pocket"
	"golang.org/x/net/html"
)

func ExampleDocument() {
	d := NewDocument(pocket.Item{})
	err := d.Fill(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	err = d.Write(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}

func Test_getURL(t *testing.T) {
	type args struct {
		attr []html.Attribute
	}
	tests := []struct {
		name         string
		args         args
		wantSource   string
		wantFilename string
		wantErr      bool
	}{
		{
			"srcset valid",
			args{
				[]html.Attribute{
					{
						Key: "src",
						Val: "http://example.com/bla",
					},
					{
						Key: "srcset",
						Val: "https://miro.medium.com/max/552/1*0oRgKqKDnBpNAMipC7juXQ.png 276w, https://miro.medium.com/max/828/1*0oRgKqKDnBpNAMipC7juXQ_big.png 414w",
					},
				},
			},
			"https://miro.medium.com/max/828/1*0oRgKqKDnBpNAMipC7juXQ_big.png",
			"1*0oRgKqKDnBpNAMipC7juXQ_big.png",
			false,
		},
		// TODO: Add test cases.
		// [{ role presentation} { src https://miro.medium.com/max/634/1*xyNSnedMmBZ-bT9vRR-bJQ.png} { width 317} { height 371} { srcset https://miro.medium.com/max/552/1*xyNSnedMmBZ-bT9vRR-bJQ.png 276w, https://miro.medium.com/max/634/1*xyNSnedMmBZ-bT9vRR-bJQ.png 317w} { sizes 317px}]
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSource, _, err := (&Document{}).getURL(tt.args.attr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSource != tt.wantSource {
				t.Errorf("getURL() gotSource = %v, want %v", gotSource, tt.wantSource)
			}
		})
	}
}
