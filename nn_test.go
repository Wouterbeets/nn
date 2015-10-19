package nn

import (
	"fmt"
	"testing"
)

func TestNewNet(test *testing.T) {
	var tests = []struct {
		inputNeurons  int
		hiddenNeurons int
		hiddenLayers  int
		outputNeurons int
		in            []float64
	}{
		{
			inputNeurons:  2,
			hiddenNeurons: 10,
			hiddenLayers:  2,
			outputNeurons: 1,
			in:            []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
		{
			inputNeurons:  19 * 19,
			hiddenNeurons: 30,
			hiddenLayers:  2,
			outputNeurons: 1,
		},
	}
	for _, t := range tests {
		n := NewNet(t.inputNeurons, t.hiddenNeurons, t.hiddenLayers, t.outputNeurons)
		n.Activate()
		n.In(t.in)
		out := n.Out()
		fmt.Println(out)
	}

}
