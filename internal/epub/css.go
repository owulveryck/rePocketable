package epub

import (
	"io"
	"io/ioutil"
)

func (d *Document) setCSS() (string, error) {
	if d.CSS == nil {
		return "", nil
	}
	file, err := ioutil.TempFile("", "mystyle*.css")
	if err != nil {
		return "", err
	}
	defer file.Close()
	//defer os.Remove(file.Name())
	_, err = io.Copy(file, d.CSS)
	if err != nil {
		return "", err
	}
	return d.Epub.AddCSS(file.Name(), "")
}
