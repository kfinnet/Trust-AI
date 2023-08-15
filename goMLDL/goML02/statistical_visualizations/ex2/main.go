package main

import (
	"log"
	"os"

	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"github.com/go-gota/gota/dataframe"
)

func main() {

	// CSV 파일을 연다
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV 파일에서 데이터프레임을 생성한다.
	irisDF := dataframe.ReadCSV(irisFile)

	// 도표를 생성하고 제목과 축 레이블을 설정한다.
	p := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "Box plots"
	p.Y.Label.Text = "Values"

	// 데이터에 대한 박스를 생성한다.
	w := vg.Points(50)

	// 데이터 집합에 있는 각 숫자 열에서 박스 도표를 생성한다
	for idx, colName := range irisDF.Names() {

		// 특정 열이 숫자 열인 경우,
		// 해당 값의 박스 도표를 생성한다.
		if colName != "species" {

			// plotter.Values값을 생성하고 데이터프레임에서 가각에
			// 해당하는 값들로 plotter.Values값을 채운다
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// 도표에 데이터를 추가한다.
			b, err := plotter.NewBoxPlot(w, float64(idx), v)
			if err != nil {
				log.Fatal(err)
			}
			p.Add(b)
		}
	}

	// x=0, x=1 등의 지정된 이름을 사용해
	// 도표의 x 축의 이름을 설정한다.
	p.NominalX("sepal_length", "sepal_width", "petal_length", "petal_width")

	if err := p.Save(6*vg.Inch, 8*vg.Inch, "boxplots.png"); err != nil {
		log.Fatal(err)
	}
}
