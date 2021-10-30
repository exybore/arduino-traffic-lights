package main

import (
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

var (
	orange = color.RGBA{R: 244, G: 122, B: 16, A: 255}
	blue = color.RGBA{R: 31, G: 133, B: 222, A: 255}
)

func plotFunction(p *plot.Plot, f func(float64) float64, xMin, xMax, yMin, yMax float64) {
	plotFunc := plotter.NewFunction(f)
	plotFunc.Color = blue

	p.Add(plotFunc)
	p.X.Min = xMin
	p.X.Max = xMax
	p.Y.Min = yMin
	p.Y.Max = yMax
}


func plotScatter(p *plot.Plot, data plotter.XYs, xLabel, yLabel string, style color.Color, line bool) {
	p.X.Label.Text = xLabel
	p.Y.Label.Text = yLabel
	p.Add(plotter.NewGrid())

	if line {
		line, err := plotter.NewLine(data)
		if err != nil {
			panic(err)
		}
		line.LineStyle.Color = style
		p.Add(line)
	} else {
		scatter, err := plotter.NewScatter(data)
		if err != nil {
			panic(err)
		}
		scatter.GlyphStyle.Color = style
		scatter.GlyphStyle.Shape = draw.CircleGlyph{}
		scatter.GlyphStyle.Radius = 0.5
		p.Add(scatter)
	}
}

func save(p *plot.Plot, title string, path string) {
	p.Title.Text = title
	if err := p.Save(8*vg.Inch, 8*vg.Inch, path); err != nil {
		panic(err)
	}
}