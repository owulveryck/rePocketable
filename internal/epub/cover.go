package epub

import (
	"image"
	"image/color"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/math/fixed"
)

func imageToCover(src, title, author, origin string) (string, error) {
	height := 2500
	width := 1600
	var img image.Image
	img = imaging.New(
		width,
		height,
		color.RGBA{128, 128, 128, 255},
	)
	resp, err := http.Get(src)
	if err == nil {
		defer resp.Body.Close()
		img, _, err = image.Decode(resp.Body)
		if err != nil {
			return "", err
		}

		img = imaging.Resize(img, width, 0, imaging.Linear)
		// put image in center of screen
		background := imaging.New(
			width,
			height,
			color.RGBA{128, 128, 128, 255},
		)
		img = imaging.PasteCenter(background, img)
	}

	addLabel(img, 62, 142, title)
	addLabel(img, 842, height-250, author)
	addLabel(img, 62, height-250, origin)
	file, err := ioutil.TempFile("", "cover*.jpg")
	if err != nil {
		return "", err
	}
	defer file.Close()
	jpeg.Encode(file, img, nil)
	return file.Name(), nil
}

func addLabel(i image.Image, x, y int, label string) error {
	labels := strings.Fields(label)
	fnt, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return err
	}
	face := truetype.NewFace(fnt, &truetype.Options{
		Size: 64,
		DPI:  110,
	})
	if img, ok := i.(*image.NRGBA); ok {
		// img is now an *image.RGBA
		col := color.RGBA{255, 255, 255, 255}
		line := 0
		var current fixed.Int26_6
		for _, label := range labels {
			label = label + " "
			if current.Round()+x > 1100 {
				current = 0
				line++
			}
			d := &font.Drawer{
				Dst:  img,
				Src:  image.NewUniform(col),
				Face: face, //basicfont.Face7x13,
				Dot:  fixed.Point26_6{fixed.Int26_6(x*64) + current, fixed.Int26_6((y + line*144) * 64)},
			}
			d.DrawString(label)
			current = current + d.MeasureString(label)
		}
	}
	return nil
}
