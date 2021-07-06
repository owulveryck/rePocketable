package db

import (
	"encoding/gob"
	"io"
	"sync"

	"github.com/owulveryck/rePocketable/internal/pocket"
)

type Database struct {
	m *sync.Map
}

func (d Database) Write(w io.Writer) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(d.m)
}

func (d *Database) Read(r io.Reader) error {
	dec := gob.NewDecoder(r)
	var m *sync.Map
	return dec.Decode(d)
}

// Store sets the value for a key.
func (d *Database) Store(key string, value pocket.Item) {
	d.Store(key, value)
}
