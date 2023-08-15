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

func main() {

	// 운전자 데이터 집합 파일을 연다
	driverDataFile, err := os.Open("fleet_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer driverDataFile.Close()

	// CSV 파일로부터 dataframe을 생성한다.
	driverDF := dataframe.ReadCSV(driverDataFile)

	// Use the Describe method to calculate summary statistics
	// for all of the columns in one shot.
	driverSummary := driverDF.Describe()

	// Output the summary statistics to stdout.
	fmt.Println(driverSummary)

	// Create a histogram for each of the columns in the dataset.
	// plotVals 변수에 데이터를 저장한다.
	for _, colName := range driverDF.Names() {

		// Create a plotter.Values value and fill it with the
		// values from the respective column of the dataframe.
		// PlotVals 변수는 도표에 대한 값을 저장하는데 사용된다.
		plotVals := make(plotter.Values, driverDF.Nrow())
		for i, floatVal := range driverDF.Col(colName).Float() {
			plotVals[i] = floatVal
		}

		// 도표를 생성한다.
		p := plot.New()
		if err != nil {
			log.Fatal(err)
		}
		p.Title.Text = fmt.Sprintf("Histogram of %s", colName)

		// Create a histogram of our values.
		h, err := plotter.NewHist(plotVals, 16)
		if err != nil {
			log.Fatal(err)
		}

		// Normalize the histogram.
		h.Normalize(1)

		// 도표를 PNG 파일로 저장한다.
		p.Add(h)
		if err := p.Save(4*vg.Inch, 4*vg.Inch, colName+"_hist.png"); err != nil {
			log.Fatal(err)
		}
	}
}
