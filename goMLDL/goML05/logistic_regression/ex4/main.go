package main

import (
	"fmt"
	"log"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// CSV 파일을 연다
	loanDataFile, err := os.Open("clean_loan_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer loanDataFile.Close()

	// CSV 파일로부터 dataframe을 생성한다.
	loanDF := dataframe.ReadCSV(loanDataFile)

	// Describe 메소드를 사용해 모든 열에 대한
	// 요약 통계를 한번에 계산한다.
	loanSummary := loanDF.Describe()

	// 요약 통계를 표준출력을 통해 출력한다.
	fmt.Println(loanSummary)

	// 데이터 집합의 모든 열에 대한 히스토그램을 생성한다.
	for _, colName := range loanDF.Names() {

		// plotter.Values 값을 생성하고 dataframe에서
		// 해당하는 값으로 plotter.Values 를 채운다.
		plotVals := make(plotter.Values, loanDF.Nrow())
		for i, floatVal := range loanDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// 도표를 만들고 제목을 설정한다.
		p := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// 원하는 값에 대한 히스토그램을 생성한다.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// 히스토그램을 정규화 한다.
		h.Normalize(1)

		// 히스토그램을 도표에 추가한다.
		p.Add(h)

		// 도표를 PNG 파일로 저장한다.
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
