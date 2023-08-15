package main

import (
	"image/color"
	"log"
	"os"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// 광고 데이터 집합 파일을 연다.
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV 파일로부터 dataframe을 생성한다
	advertDF := dataframe.ReadCSV(f)

	// 목표하는 열을 추출한다.
	yVals := advertDF.Col("Sales").Float()

	// 데이터 집합의 각 수치에 대한 신점도(scatter plot) 생성한다
	for _, colName := range advertDF.Names() {

		// pts 변수는 도표에 대한 값을 저장한다
		pts := make(plotter.XYs, advertDF.Nrow())

		// pts 변수를 데이터로 채운다
		for i, floatVal := range advertDF.Col(colName).Float() {
			pts[i].X = floatVal
			pts[i].Y = yVals[i]
		}

		// 도표를 생성한다.
		p := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.X.Label.Text = colName
		p.Y.Label.Text = "y"
		p.Add(plotter.NewGrid())

		s, err := plotter.NewScatter(pts)
		if err != nil {
			log.Fatal(err)
		}
		s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
		s.GlyphStyle.Radius = vg.Points(3)

		// 도표를 PNG파일로 저장한다
		p.Add(s)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_scatter.png"); err != nil {
			log.Fatal(err)
		}
	}
}
