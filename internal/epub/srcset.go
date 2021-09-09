package epub

import (
	"errors"
	"regexp"
	"strings"

	"strconv"
)

type srcSetElements []srcSetElement

func newSrcSetElementsFromStrings(s []string) (srcSetElements, error) {
	elements := make([]srcSetElement, 0)
	for _, s := range s {
		element, err := newSrcSetElementFromString(s)
		if err != nil {
			return srcSetElements{}, err
		}
		elements = append(elements, element)
	}
	return srcSetElements(elements), nil
}

type srcSetElement struct {
	url            string
	intrinsicWidth string
}

func newSrcSetElementFromString(s string) (srcSetElement, error) {
	elements := strings.Fields(s)
	if len(elements) != 2 {
		return srcSetElement{}, errors.New("bad set" + s)
	}
	return srcSetElement{
		url:            elements[0],
		intrinsicWidth: elements[1],
	}, nil
}

// Len is the number of elements in the collection.
func (s srcSetElements) Len() int {
	return len(s)
}

// Less reports whether the element with index i
// must sort before the element with index j.
//
// If both Less(i, j) and Less(j, i) are false,
// then the elements at index i and j are considered equal.
// Sort may place equal elements in any order in the final result,
// while Stable preserves the original input order of equal elements.
//
// Less must describe a transitive ordering:
//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//
// Note that floating-point comparison (the < operator on float32 or float64 values)
// is not a transitive ordering when not-a-number (NaN) values are involved.
// See Float64Slice.Less for a correct implementation for floating-point values.
func (s srcSetElements) Less(i int, j int) bool {
	re := regexp.MustCompile(`\d+`)
	wi, err := strconv.ParseInt(re.FindString(s[i].intrinsicWidth), 10, 32)
	if err != nil {
		panic(err)
	}
	wj, err := strconv.ParseInt(re.FindString(s[j].intrinsicWidth), 10, 32)
	if err != nil {
		panic(err)
	}
	return wi > wj
}

// Swap swaps the elements with indexes i and j.
func (s srcSetElements) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}
