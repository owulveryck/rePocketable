// Copyright ©2020 The go-latex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package epub

import (
	"log"

	"github.com/go-fonts/liberation/liberationserifbold"
	"github.com/go-fonts/liberation/liberationserifbolditalic"
	"github.com/go-fonts/liberation/liberationserifitalic"
	"github.com/go-fonts/liberation/liberationserifregular"
	"golang.org/x/image/font/sfnt"

	"github.com/go-latex/latex/font/ttf"
)

func liberationFonts() *ttf.Fonts {
	rm, err := sfnt.Parse(liberationserifregular.TTF)
	if err != nil {
		log.Fatalf("could not parse fonts: %+v", err)
	}

	it, err := sfnt.Parse(liberationserifitalic.TTF)
	if err != nil {
		log.Fatalf("could not parse fonts: %+v", err)
	}

	bf, err := sfnt.Parse(liberationserifbold.TTF)
	if err != nil {
		log.Fatalf("could not parse fonts: %+v", err)
	}

	bfit, err := sfnt.Parse(liberationserifbolditalic.TTF)
	if err != nil {
		log.Fatalf("could not parse fonts: %+v", err)
	}

	return &ttf.Fonts{
		Default: rm,
		Rm:      rm,
		It:      it,
		Bf:      bf,
		BfIt:    bfit,
	}
}
