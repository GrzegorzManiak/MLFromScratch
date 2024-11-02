package main

import (
	"MLFromScratch/internal"
)

func main() {

	inputLayer := internal.Layer{
		Name:    "Input Layer",
		Neurons: 5,
		Neuron:  internal.InputNeuron{},
	}

	hiddenLayer1 := internal.Layer{
		Name:    "Hidden Layer 1",
		Neurons: 16,
		Neuron:  internal.ReLUNeuron{},
	}

	hiddenLayer2 := internal.Layer{
		Name:    "Hidden Layer 2",
		Neurons: 16,
		Neuron:  internal.ReLUNeuron{},
	}

	outputLayer := internal.Layer{
		Name:    "Output Layer",
		Neurons: 10,
		Neuron:  internal.ReLUNeuron{},
	}

	inputLayer.GenerateNeurons()

	hiddenLayer1.GenerateNeurons()
	inputLayer.BindTo(&hiddenLayer1)

	hiddenLayer2.GenerateNeurons()
	hiddenLayer1.BindTo(&hiddenLayer2)

	outputLayer.GenerateNeurons()
	hiddenLayer2.BindTo(&outputLayer)
}
