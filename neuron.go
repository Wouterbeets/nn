package nn

import (
	"fmt"
	"math"
	"math/rand"
)

type neuron struct {
	input  []chan float64
	output []chan float64
	weight []float64
	bias   float64
	id     int
}

type inpMes struct {
	val float64
	id  int
}

func (n *neuron) mergeInputChannels() chan inpMes {
	ch := make(chan inpMes)
	for k, c := range n.input {
		go func(k int, c <-chan float64, ch chan<- inpMes) {
			for {
				ch <- inpMes{val: <-c, id: k}
			}
		}(k, c, ch)
	}
	return ch
}
func sigmoid(sum float64) float64 {
	return 1.0 / (1.0 + math.Exp(-sum))
}

func (n *neuron) activate() {
	lenInp := len(n.input)
	lenOut := len(n.output)
	c := n.mergeInputChannels()
	for {
		totalInpVal := float64(0)
		for i := 0; i < lenInp; i++ {
			res := <-c
			totalInpVal += res.val*n.weight[res.id] + n.bias
		}
		for i := 0; i < lenOut; i++ {
			n.output[i] <- sigmoid(totalInpVal)
		}
	}
}

func (n *neuron) String() string {
	str := ""
	for k, v := range n.weight {
		str += fmt.Sprintln("\t\tweight aplied to neur", k, "is", v)
	}
	for k, v := range n.input {
		str += fmt.Sprintln("\t\tchannel", k, "is", v)
	}
	return str
}

func newNeuron(input []chan float64, output []chan float64) *neuron {
	n := &neuron{
		input:  input,
		output: output,
		weight: make([]float64, len(input), len(input)),
	}
	for k, _ := range n.weight {
		n.weight[k] = rand.NormFloat64()
	}
	return n
}
