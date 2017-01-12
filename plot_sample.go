package main

import (
	"math/rand"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/plotutil"
	"github.com/gonum/plot/vg"
)

func main() {
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "graf"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"

	if err := plotutil.AddLinePoints(p, "テスト", point(5)); err != nil {
		panic(err)
	}

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "graf.png"); err != nil {
		panic(err)
	}

}

func point(n int) plotter.XYs {

	pts := make(plotter.XYs, n)

	for i := 0; i < 5; i++ {
		pts[i].X = float64(i)
		pts[i].Y = float64(i * i)
	}

	return pts
}
