package nn

import (
	"fmt"
)

type layer struct {
	neurons     []*neuron
	inputChans  []chan float64
	connections int
}

func newLayer(prevL *layer, numNeurons, numNextL int) *layer {
	l := &layer{
		neurons:    make([]*neuron, numNeurons),
		inputChans: make([]chan float64, numNeurons),
	}
	for i := 0; i < numNeurons; i++ {
		out := make([]chan float64, numNextL)
		for k := range out {
			out[k] = make(chan float64)
		}
		if prevL != nil {
			inp := make([]chan float64, len(prevL.neurons))
			for k, prevLNeur := range prevL.neurons {
				inp[k] = prevLNeur.output[i]
			}
			//fmt.Println("creating hidden neuron", numNeur)
			l.neurons[i] = newNeuron(inp, out, false)
			l.connections += len(l.neurons[i].weight)
		} else {
			inp := make([]chan float64, 1)
			inp[0] = make(chan float64)
			l.inputChans[i] = inp[0]
			//fmt.Println("creating input neuron", numNeur)
			l.neurons[i] = newNeuron(inp, out, true)
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

func (l *layer) getWeights() []float64 {
	w := make([]float64, 0, l.connections)

	for _, neur := range l.neurons {
		w = append(w, neur.weight...)
	}
	return w
}

func (l *layer) setWeigths(w []float64) {
	start, end := 0, 0
	for _, neur := range l.neurons {
		end += len(neur.input)
		neur.weight = w[start:end]
		start = end
	}
}
