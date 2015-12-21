package nn

import (
	"fmt"
)

type layer struct {
	neurons     []*neuron
	inputChans  []chan float64
	connections int
}

func newLayer(prevL *layer, numNeurons, numNextL int, weights []float64, wCount *int) *layer {
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
			subWeights := weights[*wCount+l.connections : *wCount+l.connections+len(inp)]
			l.neurons[i] = newNeuron(inp, out, false, subWeights)
			l.connections += len(l.neurons[i].weight)
		} else {
			inp := make([]chan float64, 1)
			inp[0] = make(chan float64)
			subWeights := make([]float64, 1)
			l.inputChans[i] = inp[0]
			//fmt.Println("creating input neuron", numNeur)
			l.neurons[i] = newNeuron(inp, out, true, subWeights)
		}
	}
	*wCount += l.connections
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

func (l *layer) getWeights() (w []float64) {
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
