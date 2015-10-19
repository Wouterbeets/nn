package nn

import (
	"fmt"
)

type Net struct {
	Layers     []*layer
	InputChan  []chan float64
	OutputChan []chan float64
}

func NewNet(inputNeurons, hiddenNeurons, hiddenLayers, outputNeurons int) *Net {
	n := &Net{
		Layers: make([]*layer, hiddenLayers+2),
	}
	n.Layers[0] = newLayer(nil, inputNeurons, hiddenNeurons)
	for i := 1; i < hiddenLayers+1; i++ {
		n.Layers[i] = newLayer(n.Layers[i-1], hiddenNeurons, hiddenNeurons)
	}
	n.Layers[hiddenLayers+1] = newLayer(n.Layers[hiddenLayers], outputNeurons, 1)
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
	ret := make([]float64, len(n.Layers[len(n.Layers)].neurons))
	for k, v := range n.Layers[len(n.Layers)-1].neurons {
		ret[k] = <-v.output[0]
	}
	return ret
}
