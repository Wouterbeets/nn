/*
	Package nn creates neural networks and gives you the tools to modify their internal weights.
*/
package nn

import (
	"fmt"
)

//Net holds the neural network. It must be created using its constructor. It spawns a goroutine for every neuron and
//can be used with its methods.
type Net struct {
	layers        []*layer
	inputChan     []chan float64
	outputChan    []chan float64
	inputNeurons  int
	hiddenNeurons int
	totalLayers   int
	outputNeurons int
}

//NewNet takes its size the nets layers as parametes, the minimum valaue for totalLayers is 3
// one for in , one for hidden, one for out.
func NewNet(inputNeurons, hiddenNeurons, totalLayers, outputNeurons int) *Net {
	n := &Net{
		layers:        make([]*layer, totalLayers),
		inputNeurons:  inputNeurons,
		hiddenNeurons: hiddenNeurons,
		totalLayers:   totalLayers,
		outputNeurons: outputNeurons,
	}
	n.layers[0] = newLayer(nil, inputNeurons, hiddenNeurons)
	for i := 1; i < totalLayers-2; i++ {
		n.layers[i] = newLayer(n.layers[i-1], hiddenNeurons, hiddenNeurons)
	}
	n.layers[totalLayers-2] = newLayer(n.layers[totalLayers-3], hiddenNeurons, outputNeurons)
	n.layers[totalLayers-1] = newLayer(n.layers[totalLayers-2], outputNeurons, 1)
	n.inputChan = n.getInputChans()
	n.activate()
	return n
}

//String makes Net implement the stringer interface
func (n *Net) String() string {
	str := ""
	for lNum, layer := range n.layers {
		str += fmt.Sprintln("layer", lNum, "holds\n", layer)
	}
	return str
}

//getInputChans is used internally in the constructer
func (n *Net) getInputChans() []chan float64 {
	return n.layers[0].inputChans
}

//activate calls layer.activate on all layers in all in Net
func (n *Net) activate() {
	for _, layer := range n.layers {
		layer.activate()
	}
}

//Sends the input for the brain into its input channels
// the result can be obtained by calling n.Out()
func (n *Net) In(inp []float64) error {
	if len(inp) == len(n.layers[0].neurons) {
		for k, v := range inp {
			n.inputChan[k] <- v
		}
		return nil
	} else {
		return fmt.Errorf("input must be same size as net.InputNeurons")
	}
}

//Net.Out returns the value from the output neurons
func (n *Net) Out() (outputValues []float64) {
	outputValues = make([]float64, len(n.layers[len(n.layers)-1].neurons))
	for k, v := range n.layers[len(n.layers)-1].neurons {
		outputValues[k] = <-v.output[0]
	}
	return outputValues
}

//Net.GetWeights returns  a copy of all the values correspoding to each input hannel of neach neuron
// these values contain all logic of the brain.
//these value can be used to construct copies of a neural network.
func (n *Net) GetWeights() (ret []float64) {
	size := 0
	for i, layer := range n.layers {
		if i != 0 {
			size += layer.connections
		}
	}
	ret = make([]float64, 0, size)
	for i, layer := range n.layers {
		if i != 0 {
			ret = append(ret, layer.getWeights()...)
		}
	}
	return
}

//Set Weight allows you to set the values corresponding to each input channel of each neuron
//after the operation weight[0] will be equal to the value of the wieght of the first channel of the first neuron in the first hidden layer
//wieght[1] will be equal the value of the weight of the second channel of said neuron. etc.
func (n *Net) SetWeights(weights []float64) {
	start, end := 0, 0
	for i, layer := range n.layers {
		if i != 0 {
			end = start + layer.connections
			layer.setWeigths(weights[start:end])
			start = end
		}
	}
}
