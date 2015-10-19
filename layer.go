package nn

import (
	"fmt"
)

type layer struct {
	neurons    []*neuron
	inputChans []chan float64
}

func newLayer(prevL *layer, numNeurons, numNextL int) *layer {
	l := &layer{
		neurons:    make([]*neuron, numNeurons),
		inputChans: make([]chan float64, numNeurons),
	}
	for i := 0; i < numNeurons; i++ {
		out := make([]chan float64, numNextL)
		if prevL != nil {
			_ = "breakpoint"
			inp := make([]chan float64, len(prevL.neurons))
			for k, prevLNeur := range prevL.neurons {
				inp[k] = prevLNeur.output[i]
			}
		} else {
			inp := make([]chan float64, 1)
			l.inputChans[i] = inp[0]
			l.neurons[i] = newNeuron(inp, out)
		}
	}
	return l
}

func (l *layer) String() string {
	str := ""
	for nNum, neur := range l.neurons {
		str += fmt.Sprintln("\tneur", nNum, "holds\n", neur)
	}
	return str
}

func (l *layer) activate() {
	for _, neuron := range l.neurons {
		go neuron.activate()
	}
}
