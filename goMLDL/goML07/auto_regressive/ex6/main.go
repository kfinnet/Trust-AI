package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// 로그 변환 및 차분된 데이터 집합 파일을 연다
	transFile, err := os.Open("log_diff_series.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer transFile.Close()

	// 열린 파일을 읽는 CSV reader를 생성한다.
	transReader := csv.NewReader(transFile)

	// CSV 레코드를 모두 읽는다.
	transReader.FieldsPerRecord = 2
	transData, err := transReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// 루프를 통해 데이터를 읽고 변환된 데이터를 기반으로
	// 예측을 수행한다.
	var transPredictions []float64
	for i, _ := range transData {

		// 헤더 및 처음 두 관찰은 건너 뛴다.
		// (예측을 위해 1주기 및 2주기 이전의 값이 필요하기 때문에_
		if i == 0 || i == 1 || i == 2 {
			continue
		}

		// 1주기 이전의 값을 구문 분석을 통해 읽는다.
		lagOne, err := strconv.ParseFloat(transData[i-1][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 2주기 이전의 값을 구문 분석을 통해 읽는다.
		lagTwo, err := strconv.ParseFloat(transData[i-2][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 학습을 거친 AR model을 활용해 변환된 변수를 예측한다.
		transPredictions = append(transPredictions, 0.008159+0.234953*lagOne-0.173682*lagTwo)
	}

	// original 원본 데이터 집합 파일을 연다.
	origFile, err := os.Open("AirPassengers.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer origFile.Close()

	// 열린 파일을 읽는 CSV reader를 생성한다.
	origReader := csv.NewReader(origFile)

	// CSV 레코드를 모두 읽는다.
	origReader.FieldsPerRecord = 2
	origData, err := origReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// ptsObs, ptsPred 변수는 도표에 사용되는 값을 저장하는 데 사용된다. 
	ptsObs := make(plotter.XYs, len(transPredictions))
	ptsPred := make(plotter.XYs, len(transPredictions))

	// 루프를 통해 원래대로 돌려놓은 변환을 수행하고 MAE를 계산한다.
	var mAE float64
	var cumSum float64
	for i := 4; i <= len(origData)-1; i++ {

		// 원본 관찰값을 구문 분석을 통해 읽는다.
		observed, err := strconv.ParseFloat(origData[i][1], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 원본 날짜(시간)값을 구문 분석을 통해 읽는다.
		date, err := strconv.ParseFloat(origData[i][0], 64)
		if err != nil {
			log.Fatal(err)
		}

		// 변환된 데이터를 기반으로 예측한 값의 인덱스를 구하기 위해
		// 값을 누적시킨다.
		cumSum += transPredictions[i-4]

		// 예측을 수행한다.
		predicted := math.Exp(math.Log(observed) + cumSum)

		// MAE를 누적시킨다.
		mAE += math.Abs(observed-predicted) / float64(len(transPredictions))

		// 도표를 그리기 위한 요소를 저장한다.
		ptsObs[i-4].X = date
		ptsPred[i-4].X = date
		ptsObs[i-4].Y = observed
		ptsPred[i-4].Y = predicted
	}

	// 표준출력을 통해 MAE 를 출력한다.
	fmt.Printf("\nMAE = %0.2f\n\n", mAE)

	// 도표를 생성한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	p.X.Label.Text = "time"
	p.Y.Label.Text = "passengers"
	p.Add(plotter.NewGrid())

	// 시계열 데이터에 대한 도표 지점들을 추가한다.
	lObs, err := plotter.NewLine(ptsObs)
	if err != nil {
		log.Fatal(err)
	}
	lObs.LineStyle.Width = vg.Points(1)

	lPred, err := plotter.NewLine(ptsPred)
	if err != nil {
		log.Fatal(err)
	}
	lPred.LineStyle.Width = vg.Points(1)
	lPred.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}

	// 도표를 PNG 파일에 저장한다.
	p.Add(lObs, lPred)
	p.Legend.Add("Observed", lObs)
	p.Legend.Add("Predicted", lPred)
	if err := p.Save(10*vg.Inch, 4*vg.Inch, "passengers_ts.png"); err != nil {
		log.Fatal(err)
	}
}
