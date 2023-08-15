package main

import (
	"fmt"
	"log"
	"os"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot"
	"github.com/go-gota/gota/dataframe"
)
// 히스토그램을 정규화(h.Normalize()를 사용해서)한 것에 주의한다.
// 개수가 다른 분포와 비교해야 하는 경우가 발생할 수 있기 때문에
// 정규화 하는것이 일반적이다.
// 히스토그램을 정규화하면 다른 분포를 나란히 표시할 수 있다.
func main() {

	// CSV 파일 열기
	irisFile, err := os.Open("../data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	// CSV 파일에서 데이터프레임 생성하기
	irisDF := dataframe.ReadCSV(irisFile)

	// 데이터 집합에 있는 각 숫자 열에서 히스토그램을 생성한다.
	for _, colName := range irisDF.Names() {

		// 특정 열이 숫자 열인 경우
		// 해당 값의 히스토그램을 생성한다.
		if colName != "species" {

			// plotter.Values값을 생성하고 데이터프레임에서 각각에
			// 해당하는 값들로 plotter.Values값을 채운다.
			v := make(plotter.Values, irisDF.Nrow())
			for i, floatVal := range irisDF.Col(colName).Float() {
				v[i] = floatVal
			}

			// 도표를 만들고 제목을 설정한다.
			//p, err := plot.New()
			p := plot.New()
			if err != nil {
				log.Fatal(err)
			}
			p.Title.Text = fmt.Sprintf("Histogram of a %s", colName)

		    // 표준 정규 분포로 그려지는 히스토그램을 만든다.
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				log.Fatal(err)
			}

			// 막대그래프를 정규화한다.
			h.Normalize(1)

			// 히스토그램을 도표에 추가한다.
			p.Add(h)

			// 도표를 PNG 파일로 저장한다.
			if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
				log.Fatal(err)
			}
		}
	}
}
// 이 분포는 다른 분포들과 서로 다른 모양을 가지고 있다.
// sepal_width 분포는 종형곡선 Bell Curve 또는 정규/가우시안 분포와 유사
// 반면에 petal 분포는 서로 다른 갓들이 두 분류로 나뉜 것처럼 보인다
// 나중에 이런 관찰 방법을 활용해 머신 러닝 모델을 개발할 예정이다.
// 시각화가 데이터에 대한 멘탈 모델 Mental Model 개발하는데 
//  어떻게 도움되는지 주목해 살펴보자