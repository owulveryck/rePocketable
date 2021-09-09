package epub

import (
	"sort"
	"testing"
)

func Test_newSrcSetElementsFromStrings(t *testing.T) {
	set, err := newSrcSetElementsFromStrings([]string{
		"a 123w",
		"b 24w",
		"ok 1245w",
		"ok 2x",
	})
	if err != nil {
		t.Fatal(err)
	}
	sort.Sort(set)
	if set[0].url != "ok" {
		t.Fatal(set)
	}
}
