package main

import (
	"image/color"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// CSV 파일을 연다.
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// CSV 파일로부터 dataframe 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// 승객의 수에 해당하는 열(AirPassengers)에서 데이터를 추출한다.
	yVals := passengersDF.Col("AirPassengers").Float()

	// pts 변수는 도표에 사용될 값을 저장하는 데 사용된다.
	pts := make(plotter.XYs, passengersDF.Nrow())

	// pts 변수를 데이터로 채운다
	for i, floatVal := range passengersDF.Col("time").Float() {
		pts[i].X = floatVal
		pts[i].Y = yVals[i]
	}

	// 도표를 생성한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	// 시계열에 대한 직선 도표의 위치들을 추가한다.
	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// 도표를 PNG 파일로 저장한다.
	p.Add(l)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}
}
