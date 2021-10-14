package epub

import (
	"io/ioutil"

	"github.com/gofrs/uuid"
	"github.com/vincent-petithory/dataurl"
)

func (d *Document) setCSS() (string, error) {
	content, err := ioutil.ReadAll(d.CSS)
	if err != nil {
		return "", err
	}
	return d.Epub.AddCSS(dataurl.EncodeBytes(content), uuid.Must(uuid.NewV4()).String()+".css")
}
