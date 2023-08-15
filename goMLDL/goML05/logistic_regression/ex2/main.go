package main

import (
	//"errors"
	"image/color"
	"log"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// 새도표를 생성한다
	p := plot.New()  //
	//if err != nil {
	//	log.Fatal(err)
	//}
	p.Title.Text = "Logistic Function"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "f(x)"

	// plotter 함수를 생성한다
	logisticPlotter := plotter.NewFunction(func(x float64) float64 { return logistic(x) })
	logisticPlotter.Color = color.RGBA{B: 255, A: 255}

	// plotter 함수를 도표에 추가한다.
	p.Add(logisticPlotter)

	// 축의 범위를 설정한다. 다른 데이터 집합과는 달리,
	// 함수는 축의 번위를 자동으로 설정하지 않기 때문에
	// 함수는 x와 y 값의 유한한 범위를 가져야 할 필요는 없다.
	p.X.Min = -10
	p.X.Max = 10
	p.Y.Min = -0.1
	p.Y.Max = 1.1

	// 도표를 PNG 파일로 저장한다.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "logistic.png"); err != nil {
		log.Fatal(err)
	}
}

// logistic 함수는 로지스틱 함수를 구현하며
// 로지스틱 회귀분석에 사용된다.
func logistic(x float64) float64 {
	return 1 / (1 + math.Exp(-x))
}
