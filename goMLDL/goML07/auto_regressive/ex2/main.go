package main

import (
	"encoding/csv"
	"image/color"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// CSV 파일을 연다
	passengersFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer passengersFile.Close()

	// CSV파일로 부터 datafream을 생성한다.
	passengersDF := dataframe.ReadCSV(passengersFile)

	// AirPassengers 열을 float 슬라이스(동적배열 포인트타입)로 읽는다
	passengerVals := passengersDF.Col("AirPassengers").Float()
	timeVals := passengersDF.Col("time").Float()

	// pts 변수는 도표에 사용되는 변수를 저장한다.
	pts := make(plotter.XYs, passengersDF.Nrow()-1)

	// differenced 변수는 새로운 CSV 파일에 저장될
	// 변환된 값을 저장하는 데 사용된다.
	var differenced [][]string
	differenced = append(differenced, []string{"time", "log_differenced_passengers"})

	// pts 변수에 값을 채운다.
	for i := 1; i < len(passengerVals); i++ {
		pts[i-1].X = timeVals[i]
		pts[i-1].Y = math.Log(passengerVals[i]) - math.Log(passengerVals[i-1])
		differenced = append(differenced, []string{
			strconv.FormatFloat(timeVals[i], 'f', -1, 64),
			strconv.FormatFloat(math.Log(passengerVals[i])-math.Log(passengerVals[i-1]), 'f', -1, 64),
		})
	}

	// 도표를 생성한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "log(differenced passengers)"
	p.Add(plotter.NewGrid())

	// 시계열에 대한 직선 도표 지점을 추가한다.
	l, err := plotter.NewLine(pts)
	if err != nil {
		log.Fatal(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// 도표를 PNG 파일로 저장한다.
	p.Add(l)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "log_diff_passengers_ts.png"); err != nil {
		log.Fatal(err)
	}

	// 변환된 데이터를 새로운 CSV 에 저장한다.
	f, err := os.Create("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(differenced)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
