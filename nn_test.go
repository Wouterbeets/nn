package nn

import (
	//"fmt"
	"testing"
)

func TestNewNet(test *testing.T) {
	var tests = []struct {
		inputNeurons  int
		hiddenNeurons int
		totalLayers   int
		outputNeurons int
		in            []float64
	}{

		{
			inputNeurons:  2,
			hiddenNeurons: 3,
			totalLayers:   4,
			outputNeurons: 1,
			in:            []float64{0.99, 0.1111},
		},

		{
			inputNeurons:  3,
			hiddenNeurons: 4,
			totalLayers:   5,
			outputNeurons: 2,
			in:            []float64{0.99, 0.1111, 0.15},
		},

		{
			inputNeurons:  4,
			hiddenNeurons: 15,
			totalLayers:   16,
			outputNeurons: 3,
			in:            []float64{0.99, 0.1111, 0.15, 12.2},
		},

		{
			inputNeurons:  4,
			hiddenNeurons: 23,
			totalLayers:   26,
			outputNeurons: 3,
			in:            []float64{0.99, 0.1111, 0.15, 12.2},
		},

		{
			inputNeurons:  4,
			hiddenNeurons: 1,
			totalLayers:   3,
			outputNeurons: 3,
			in:            []float64{0.99, 0.1111, 0.15, 12.2},
		},

		{
			inputNeurons:  4,
			hiddenNeurons: 3,
			totalLayers:   3,
			outputNeurons: 3,
			in:            []float64{0.99, 0.1111, 0.15, 12.2},
		},
	}

	for _, t := range tests {
		n := NewNet(t.inputNeurons, t.hiddenNeurons, t.totalLayers, t.outputNeurons)
		w := n.GetWeights()
		//fmt.Println("\nweights are", n.GetWeights())
		for i, _ := range w {
			w[i] = float64(i)
		}
		//fmt.Println("\nweights are", n.GetWeights())
		n.In(t.in)
		_ = n.Out()
	}
}

func BenchmarkSetWeights(b *testing.B) {
	n := NewNet(4, 60, 10, 1)
	for i := 0; i < b.N; i++ {
		w := n.GetWeights()
		n.SetWeights(w)
	}
}
