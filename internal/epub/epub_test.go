package epub

import (
	"context"
	"log"
	"os"

	"github.com/owulveryck/rePocketable/internal/pocket"
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
