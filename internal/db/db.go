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

func NewDatabase() *Database {
	return &Database{
		m: &sync.Map{},
	}
}

func (d Database) Write(w io.Writer) error {
	bkp := make(map[int]pocket.Item)
	d.m.Range(func(key, value interface{}) bool {
		k, ok := key.(int)
		if !ok {
			return false
		}
		v, ok := value.(pocket.Item)
		if !ok {
			return false
		}
		bkp[k] = v
		return true
	})
	enc := gob.NewEncoder(w)
	return enc.Encode(bkp)
}

func (d *Database) Read(r io.Reader) error {
	dec := gob.NewDecoder(r)
	var bkp map[int]pocket.Item
	err := dec.Decode(&bkp)
	if err != nil {
		return err
	}
	for k, v := range bkp {
		d.Store(k, v)
	}
	return nil
}

// Store sets the value for a key.
func (d *Database) Store(key int, value pocket.Item) {
	d.m.Store(key, value)
}

func (d *Database) Load(key int) (value pocket.Item, ok bool) {
	v, ok := d.m.Load(key)
	if !ok {
		return pocket.Item{}, ok
	}
	item, ok := v.(pocket.Item)
	return item, ok
}

func (d *Database) Range(f func(key interface{}, value interface{}) bool) {
	d.m.Range(f)
}
