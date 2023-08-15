package main

import (
	"image/color"
	"log"
	"os"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// 운전자 데이터 집합 파일을 연다
	f, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV 파일로부터 dataframe 을 생성한다.
	driverDF := dataframe.ReadCSV(f)

	// Distance_Feature에 대한 열에서 데이터를 추출한다.
	yVals := driverDF.Col("Distance_Feature").Float()

	// pts 변수는 도표에 대한 값을 저장하는 데 사용된다.
	pts := make(plotter.XYs, driverDF.Nrow())

	// pts 변수에 데이터를 저장한다.
	for i, floatVal := range driverDF.Col("Speeding_Feature").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// 도표를 생성한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "Speeding"
	p.Y.Label.Text = "Distance"
	p.Add(plotter.NewGrid())

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// 도표를 PNG 파일로 저장한다.
	p.Add(s)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "fleet_data_scatter.png"); err != nil {
		log.Fatal(err)
	}
}
