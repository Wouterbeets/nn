package nn

import (
	"fmt"
)

const (
	SIZE = 20
)

type Net struct {
	Layers        []*layer
	InputChan     []chan float64
	OutputChan    []chan float64
	inputNeurons  int
	hiddenNeurons int
	totalLayers   int
	outputNeurons int
}

func NewNet(inputNeurons, hiddenNeurons, totalLayers, outputNeurons int) *Net {
	n := &Net{
		Layers:        make([]*layer, totalLayers),
		inputNeurons:  inputNeurons,
		hiddenNeurons: hiddenNeurons,
		totalLayers:   totalLayers,
		outputNeurons: outputNeurons,
	}
	n.Layers[0] = newLayer(nil, inputNeurons, hiddenNeurons)
	for i := 1; i < totalLayers-2; i++ {
		//fmt.Println("hidden neurons:")
		n.Layers[i] = newLayer(n.Layers[i-1], hiddenNeurons, hiddenNeurons)
	}
	n.Layers[totalLayers-2] = newLayer(n.Layers[totalLayers-3], hiddenNeurons, outputNeurons)
	//fmt.Println("output neurons:")
	n.Layers[totalLayers-1] = newLayer(n.Layers[totalLayers-2], outputNeurons, 1)
	n.InputChan = n.getInputChans()
	return n
}

func (n *Net) String() string {
	str := ""
	for lNum, layer := range n.Layers {
		str += fmt.Sprintln("layer", lNum, "holds\n", layer)
	}
	return str
}

func (n *Net) getInputChans() []chan float64 {
	return n.Layers[0].inputChans
}

func (n *Net) Activate() {
	for _, layer := range n.Layers {
		layer.activate()
	}
}

func (n *Net) In(inp []float64) {
	if len(inp) == len(n.Layers[0].neurons) {
		for k, v := range inp {
			n.InputChan[k] <- v
		}
	}
}

func (n *Net) Out() []float64 {
	ret := make([]float64, len(n.Layers[len(n.Layers)-1].neurons))
	for k, v := range n.Layers[len(n.Layers)-1].neurons {
		ret[k] = <-v.output[0]
	}
	return ret
}

func (n *Net) GetWeights() (ret []float64) {
	size := 0
	for i, layer := range n.Layers {
		if i != 0 {
			size += layer.connections
		}
	}
	ret = make([]float64, 0, size)
	_ = "breakpoint"
	for i, layer := range n.Layers {
		if i != 0 {
			ret = append(ret, layer.getWeights()...)
		}
	}
	return
}

func (n *Net) SetWeights(weights []float64) {
	start, end := 0, 0
	for i, layer := range n.Layers {
		if i != 0 {
			end = start + layer.connections
			layer.setWeigths(weights[start:end])
			start = end
		}
	}
}
