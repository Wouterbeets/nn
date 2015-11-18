package nn

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type neuron struct {
	input   []chan float64
	output  []chan float64
	weight  []float64
	bias    float64
	id      int
	isInput bool
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
	//fmt.Println("neuron", n.id, "activated")
	lenOut := len(n.output)
	lenIn := len(n.input)
	c := n.mergeInputChannels()
	for {
		totalInpVal := float64(0)
		for i := 0; i < lenIn; i++ {
			if n.id == 7 {
				_ = "breakpoint"
			}
			res := <-c
			totalInpVal += res.val * n.weight[res.id]
			//fmt.Println("neur", n.id, "received", res, "weight", n.weight[res.id])
		}
		for i := 0; i < lenOut; i++ {
			//fmt.Println("sending from ", n.id, "val", totalInpVal)
			n.output[i] <- sigmoid(totalInpVal) //+ n.bias
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

var numNeur = 0

func newNeuron(input []chan float64, output []chan float64, isInput bool) *neuron {
	n := &neuron{
		input:   input,
		output:  output,
		weight:  make([]float64, len(input), len(input)),
		id:      numNeur,
		isInput: isInput,
		bias:    -0.5,
	}
	numNeur++
	if n.isInput == true {
		for k, _ := range n.weight {
			n.weight[k] = 1.0
		}
	} else {
		for k, _ := range n.weight {
			n.weight[k] = rand.NormFloat64() * 10
		}
	}
	//fmt.Println(n.weight)
	return n
}
