// Copyright ©2020 The go-latex Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mtex

import (
	"fmt"
	"math"

	"github.com/go-latex/latex/drawtex"
	"github.com/go-latex/latex/font/ttf"
	"github.com/go-latex/latex/tex"
)

type Renderer interface {
	Render(w, h, dpi float64, cnv *drawtex.Canvas) error
}

func Render(dst Renderer, expr string, size, dpi float64) error {
	cnv := drawtex.New()
	p := newParser(dst, ttf.New(cnv))
	box, err := p.parse(expr, size, 72)
	if err != nil {
		return fmt.Errorf("could not parse math expression: %w", err)
	}

	var sh tex.Ship
	sh.Call(0, 0, box.(tex.Tree))

	w := box.Width()
	h := box.Height()
	d := box.Depth()

	err = dst.Render(w/72, math.Ceil(h+math.Max(d, 0))/72, dpi, cnv)
	if err != nil {
		return fmt.Errorf("could not render math expression: %w", err)
	}

	return nil
}
