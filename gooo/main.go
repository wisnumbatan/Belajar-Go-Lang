package main

import (
	"fmt"
	"math"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Mendefinisikan slice x
	x := make([]float64, 200)
	for i := range x {
		x[i] = -2.0 + float64(i)*(7.0/199)
	}

	// Fungsi f(x) = x^2 + 1
	f := func(x float64) float64 {
		return math.Pow(x, 2) + 1
	}

	// Turunan df(x) = 2x
	df := func(x float64) float64 {
		return 2 * x
	}

	// Hitung f(x) dan df(x)
	fx := make(plotter.XYs, len(x))
	dfx := make(plotter.XYs, len(x))
	dfxFromDiff := make(plotter.XYs, len(x)-1)

	for i, val := range x {
		fx[i].X = val
		fx[i].Y = f(val)

		dfx[i].X = val
		dfx[i].Y = df(val)
	}

	floats.Span(dfxFromDiff, 0, 7, func(i float64) float64 {
		return (f(i+0.01) - f(i-0.01)) / 0.02
	})

	// Plot grafik f(x) dan df(x)
	p, _ := plot.New()
	p.Title.Text = "Grafik f(x) dan df(x)"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x) / df(x)"

	fPlot, _ := plotter.NewLine(fx)
	fPlot.LineStyle.Width = vg.Points(1)
	fPlot.LineStyle.Color = plotutil.Color(0)
	fPlot.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	p.Add(fPlot)

	dfPlot, _ := plotter.NewLine(dfx)
	dfPlot.LineStyle.Width = vg.Points(1)
	dfPlot.LineStyle.Color = plotutil.Color(1)
	p.Add(dfPlot)

	dfDiffPlot, _ := plotter.NewLine(dfxFromDiff)
	dfDiffPlot.LineStyle.Width = vg.Points(1)
	dfDiffPlot.LineStyle.Color = plotutil.Color(2)
	p.Add(dfDiffPlot)

	p.Legend.Add("f(x) = x^2 + 1", fPlot)
	p.Legend.Add("df(x) = 2x", dfPlot)
	p.Legend.Add("df(x) (from diff)", dfDiffPlot)
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "grafik_f_df.png"); err != nil {
		fmt.Println(err)
	}
}
