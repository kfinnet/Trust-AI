package main

import (
	"fmt"
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

	// CSV 파일로부터 dataframe을 생성한다
	advertDF := dataframe.ReadCSV(f)

	// 데이터 집합의 각 열에 대한 히스토그램을 생성한다
	for _, colName := range advertDF.Names() {

		// plotter.Values 값을 생성하고 datafreame 각 열에 있는 값으로
		// Plotter.Values를 채운다
		plotVals := make(plotter.Values, advertDF.Nrow())
		for i, floatVal := range advertDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// 도표를 만들고 제목을 설정한다
		p := plot.New()   
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		// 표준 법선으로부터 그려지는 값의
		// 히스토그램을 생성한다
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// 히스토그램을 정규화한다.
		h.Normalize(1)

		// 히스토그램을 도표에 추가한다.
		p.Add(h)

		// 도표를 PNG파일로 저장한다
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
