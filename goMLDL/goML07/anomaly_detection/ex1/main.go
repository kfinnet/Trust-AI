package main

import (
	"fmt"
	"log"

	"github.com/lytics/anomalyzer"
)

func main() {

	// 어떤 이상 감지 메소드를 사용할 지와 같은 설정을 적용해
	// AnomalyzerConf값을 초기화 한다.
	conf := &anomalyzer.AnomalyzerConf{
		Sensitivity: 0.1,
		UpperBound:  5,
		LowerBound:  anomalyzer.NA, // ignore the lower bound
		ActiveSize:  1,
		NSeasons:    4,
		Methods:     []string{"diff", "fence", "highrank", "lowrank", "magnitude"},
	}

	// 주기적인 관찰 데이터가 포함되는 시계열 데이터를 float 슬라이스 생성한다.
	// 이 값들은 앞의 예제에서 사용했던 것처럼
	// 데이터베이스나 파일에서 읽어올 수 있다
	ts := []float64{0.1, 0.2, 0.5, 0.12, 0.38, 0.9, 0.74}

	// 기존의 시계열 데이터 값과 설정을 기반으로
	// 새 anomalyer 생성한다.
	anom, err := anomalyzer.NewAnomalyzer(conf, ts)
	if err != nil {
		log.Fatal(err)
	}

	// Anomalyzer에 새로 관찰된 값을 추가한다.
	// Anomalyzer는 시계열 데이터의 기존 값을 참조해 값을 분석하고
	// 해당 값이 비정상적일 확률을 출력한다.
	prob := anom.Push(15.2)
	fmt.Printf("Probability of 15.2 being anomalous: %0.2f\n", prob)

	prob = anom.Push(0.43)
	fmt.Printf("Probability of 0.33 being anomalous: %0.2f\n", prob)
}
