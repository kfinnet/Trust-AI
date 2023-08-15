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

	// 광고 데이터 집합 파일을 연다
	f, err := os.Open("Advertising.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// CSV파일에서 datafram을 생성한다
	advertDF := dataframe.ReadCSV(f)

	// 목표하는 열을 추출한다.
	yVals := advertDF.Col("Sales").Float()

	// pts 변수는 도표를 그리는 데 사용되는 값을 저장한다.
	pts := make(plotter.XYs, advertDF.Nrow())

	// ptsPred 변수는 도표를 그리는 데 사용되는 예측값을 저장한다.
	ptsPred := make(plotter.XYs, advertDF.Nrow())

	// pts 변수를 데이터로 채운다.
	for i, floatVal := range advertDF.Col("TV").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
		ptsPred[i].X = floatVal
		ptsPred[i].Y = predict(floatVal)
	}

	// 도표를 생성한다.
	p := plot.New()  // err 변수를 없앤다
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "TV"
	p.Y.Label.Text = "Sales"
	p.Add(plotter.NewGrid())

	// 관찰에 대한 산점도값을 추가한다.
	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.GlyphStyle.Radius = vg.Points(3)

	// 예측에 대한 직선 도표값을 추가한다.
	l, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// 도표를 PNG 파일로 저장한다.
	p.Add(s, l)
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "regression_line.png"); err != nil {
		log.Fatal(err)
	}
}

// predict TV 값을 기반으로 예측을 수행한다.
func predict(tv float64) float64 {
	return 7.07 + tv*0.05
}
