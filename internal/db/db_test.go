package db

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/owulveryck/rePocketable/internal/pocket"
)

func TestDatabase(t *testing.T) {
	date, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	originalItem := pocket.Item{
		ItemID:    42,
		TimeAdded: pocket.Time(date),
	}
	d := NewDatabase()
	d.Store(42, originalItem)
	d.Store(43, originalItem)
	var backup bytes.Buffer
	err := d.Write(&backup)
	if err != nil {
		t.Fatal(err)
	}
	d2 := NewDatabase()
	err = d2.Read(&backup)
	if err != nil {
		t.Fatal(err)
	}
	item, ok := d2.Load(42)
	if !ok {
		t.Fatal("value not found")
	}
	if !reflect.DeepEqual(originalItem, item) {
		t.Fatalf("bad value, expected %v, got %v", time.Time(originalItem.TimeAdded), time.Time(item.TimeAdded))
	}
	_, ok = d2.Load(43)
	if !ok {
		t.Fatal("value not found")
	}
}
