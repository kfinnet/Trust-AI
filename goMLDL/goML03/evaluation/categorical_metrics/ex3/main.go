package main

import (
	"fmt"

	"github.com/gonum/stat"
	"gonum.org/v1/gonum/integrate"
)

func main() {

	// 점수 및 클래스를 정의한다
	scores := []float64{0.1, 0.35, 0.4, 0.8}
	classes := []bool{true, false, true, false}

	// true positive 비율과 false positive 비율을 
	// 계산한다.
	tpr, fpr := stat.ROC(0, scores, classes, nil)

	// AUC(Area Under Curve)를 계산한다.
	auc := integrate.Trapezoidal(fpr, tpr)

	// 표준 출력을 통해 결과를 출력한다.
	fmt.Printf("x축 tpr=FP/(FP + TN)):true  positive rate: %v\n", tpr)
	fmt.Printf("y축 Fpr=TP/(TP + FN))[== Recall]:false positive rate: %v\n", fpr)
	fmt.Printf("정확도 auc(Area Under Curve): %v\n", auc)
}
