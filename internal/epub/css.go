package epub

import (
	"io/ioutil"

	"github.com/vincent-petithory/dataurl"
)

func (d *Document) setCSS() (string, error) {
	content, err := ioutil.ReadAll(d.CSS)
	if err != nil {
		return "", err
	}
	return d.Epub.AddCSS(dataurl.EncodeBytes(content), "")
}
