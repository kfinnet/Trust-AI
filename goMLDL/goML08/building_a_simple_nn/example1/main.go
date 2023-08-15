package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

// neuralNet 구조체는 훈련(학습)된 신경망을 정의하는
// 모든정보를 포함한다.
type neuralNet struct {
	config  neuralNetConfig
	wHidden *mat.Dense
	bHidden *mat.Dense
	wOut    *mat.Dense
	bOut    *mat.Dense
}

// neuralNetConfig 구조체는 신경망 아키텍쳐 및
// 학습 매개변수를 정의한다.
type neuralNetConfig struct {
	inputNeurons  int
	outputNeurons int
	hiddenNeurons int
	numEpochs     int
	learningRate  float64
}

func main() {

	// Define our input attributes.
	input := mat.NewDense(3, 4, []float64{
		1.0, 0.0, 1.0, 0.0,
		1.0, 0.0, 1.0, 1.0,
		0.0, 1.0, 0.0, 1.0,
	})

	// Define our labels.
	labels := mat.NewDense(3, 1, []float64{1.0, 1.0, 0.0})

	// Define our network architecture and
	// learning parameters.
	config := neuralNetConfig{
		inputNeurons:  4,
		outputNeurons: 1,
		hiddenNeurons: 3,
		numEpochs:     5000,
		learningRate:  0.3,
	}

	// Train the neural network.
	network := newNetwork(config)
	if err := network.train(input, labels); err != nil {
		log.Fatal(err)
	}

	// Output the weights that define our network!
	f := mat.Formatted(network.wHidden, mat.Prefix("          "))
	fmt.Printf("\nwHidden = % v\n\n", f)

	f = mat.Formatted(network.bHidden, mat.Prefix("    "))
	fmt.Printf("\nbHidden = % v\n\n", f)

	f = mat.Formatted(network.wOut, mat.Prefix("       "))
	fmt.Printf("\nwOut = % v\n\n", f)

	f = mat.Formatted(network.bOut, mat.Prefix("    "))
	fmt.Printf("\nbOut = % v\n\n", f)
}

// NewNetwork 함수는 새로운 신경망을 초기화한다.
func newNetwork(config neuralNetConfig) *neuralNet {
	return &neuralNet{config: config}
}

// Train 함수는 역전파를 사용해 신경망을 훈련(학습)시킨다.
func (nn *neuralNet) train(x, y *mat.Dense) error {

	// 바이어스/가중치를 초기화한다.
	randSource := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(randSource)

	wHiddenRaw := make([]float64, nn.config.hiddenNeurons*nn.config.inputNeurons)
	bHiddenRaw := make([]float64, nn.config.hiddenNeurons)
	wOutRaw := make([]float64, nn.config.outputNeurons*nn.config.hiddenNeurons)
	bOutRaw := make([]float64, nn.config.outputNeurons)

	for _, param := range [][]float64{wHiddenRaw, bHiddenRaw, wOutRaw, bOutRaw} {
		for i := range param {
			param[i] = randGen.Float64()
		}
	}

	wHidden := mat.NewDense(nn.config.inputNeurons, nn.config.hiddenNeurons, wHiddenRaw)
	bHidden := mat.NewDense(1, nn.config.hiddenNeurons, bHiddenRaw)
	wOut := mat.NewDense(nn.config.hiddenNeurons, nn.config.outputNeurons, wOutRaw)
	bOut := mat.NewDense(1, nn.config.outputNeurons, bOutRaw)

	// 신경망의 출력을 정의한다.
	output := mat.NewDense(0, 0, nil)

	// 루프를 통해 모델을 훈련시키기 위해
	// 역전파를 활용한 작업을 수행한다.
	for i := 0; i < nn.config.numEpochs; i++ {

		// Complete the feed forward process.
		hiddenLayerInput := mat.NewDense(0, 0, nil)
		hiddenLayerInput.Mul(x, wHidden)
		addBHidden := func(_, col int, v float64) float64 { return v + bHidden.At(0, col) }
		hiddenLayerInput.Apply(addBHidden, hiddenLayerInput)

		hiddenLayerActivations := mat.NewDense(0, 0, nil)
		applySigmoid := func(_, _ int, v float64) float64 { return sigmoid(v) }
		hiddenLayerActivations.Apply(applySigmoid, hiddenLayerInput)

		outputLayerInput := mat.NewDense(0, 0, nil)
		outputLayerInput.Mul(hiddenLayerActivations, wOut)
		addBOut := func(_, col int, v float64) float64 { return v + bOut.At(0, col) }
		outputLayerInput.Apply(addBOut, outputLayerInput)
		output.Apply(applySigmoid, outputLayerInput)

		// 역전파 작업을 완성한다.
		networkError := mat.NewDense(0, 0, nil)
		networkError.Sub(y, output)

		slopeOutputLayer := mat.NewDense(0, 0, nil)
		applySigmoidPrime := func(_, _ int, v float64) float64 { return sigmoidPrime(v) }
		slopeOutputLayer.Apply(applySigmoidPrime, output)
		slopeHiddenLayer := mat.NewDense(0, 0, nil)
		slopeHiddenLayer.Apply(applySigmoidPrime, hiddenLayerActivations)

		dOutput := mat.NewDense(0, 0, nil)
		dOutput.MulElem(networkError, slopeOutputLayer)
		errorAtHiddenLayer := mat.NewDense(0, 0, nil)
		errorAtHiddenLayer.Mul(dOutput, wOut.T())

		dHiddenLayer := mat.NewDense(0, 0, nil)
		dHiddenLayer.MulElem(errorAtHiddenLayer, slopeHiddenLayer)

		// Adjust the parameters.매개변수를 조정한다
		wOutAdj := mat.NewDense(0, 0, nil)
		wOutAdj.Mul(hiddenLayerActivations.T(), dOutput)
		wOutAdj.Scale(nn.config.learningRate, wOutAdj)
		wOut.Add(wOut, wOutAdj)

		bOutAdj, err := sumAlongAxis(0, dOutput)
		if err != nil {
			return err
		}
		bOutAdj.Scale(nn.config.learningRate, bOutAdj)
		bOut.Add(bOut, bOutAdj)

		wHiddenAdj := mat.NewDense(0, 0, nil)
		wHiddenAdj.Mul(x.T(), dHiddenLayer)
		wHiddenAdj.Scale(nn.config.learningRate, wHiddenAdj)
		wHidden.Add(wHidden, wHiddenAdj)

		bHiddenAdj, err := sumAlongAxis(0, dHiddenLayer)
		if err != nil {
			return err
		}
		bHiddenAdj.Scale(nn.config.learningRate, bHiddenAdj)
		bHidden.Add(bHidden, bHiddenAdj)
	}

	// Define our trained neural network.
	nn.wHidden = wHidden
	nn.bHidden = bHidden
	nn.wOut = wOut
	nn.bOut = bOut

	return nil
}

// sigmoid implements the sigmoid function
// for use in activation functions.
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// sigmoidPrime implements the derivative
// of the sigmoid function for backpropagation.
func sigmoidPrime(x float64) float64 {
	return x * (1.0 - x)
}

// sumAlongAxis sums a matrix along a
// particular dimension, preserving the
// other dimension.
func sumAlongAxis(axis int, m *mat.Dense) (*mat.Dense, error) {

	numRows, numCols := m.Dims()

	var output *mat.Dense

	switch axis {
	case 0:
		data := make([]float64, numCols)
		for i := 0; i < numCols; i++ {
			col := mat.Col(nil, i, m)
			data[i] = floats.Sum(col)
		}
		output = mat.NewDense(1, numCols, data)
	case 1:
		data := make([]float64, numRows)
		for i := 0; i < numRows; i++ {
			row := mat.Row(nil, i, m)
			data[i] = floats.Sum(row)
		}
		output = mat.NewDense(numRows, 1, data)
	default:
		return nil, errors.New("invalid axis, must be 0 or 1")
	}

	return output, nil
}
