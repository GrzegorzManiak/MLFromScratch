package internal

import "math/rand"

type Neuron interface {
	Calculate(axons []Axon) float64
	GetActivation() float64
	Create() Neuron
	GetID() uint64
	RandomizeBias()
}

var IDCounter uint64 = 0

type BaseNeuron struct {
	ID         uint64
	Bias       float64
	Activation float64
}

func (n BaseNeuron) GetID() uint64 {
	return n.ID
}

func (n BaseNeuron) GetActivation() float64 {
	return n.Activation
}

func (n *BaseNeuron) RandomizeBias() {
	n.Bias = rand.Float64()*2 - 1
}

// INPUT NEURON
//

type InputNeuron struct {
	BaseNeuron
}

func (n InputNeuron) Calculate(axons []Axon) float64 {
	return n.Activation
}

func (n InputNeuron) Create() Neuron {
	IDCounter++
	neuron := InputNeuron{BaseNeuron{ID: IDCounter}}
	return &neuron
}

// RELU NEURON
//

type ReLUNeuron struct {
	BaseNeuron
}

func (n ReLUNeuron) Calculate(axons []Axon) float64 {
	var weightedSum float64
	for _, axon := range axons {
		weightedSum += axon.Weight * (*axon.NeuronA).GetActivation()
	}
	weightedSum += n.Bias
	return ReLU(weightedSum)
}

func (n ReLUNeuron) Create() Neuron {
	IDCounter++
	neuron := ReLUNeuron{BaseNeuron{ID: IDCounter}}
	return &neuron
}
