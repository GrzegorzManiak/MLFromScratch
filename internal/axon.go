package internal

import "math/rand"

type Axon struct {
	Weight  float64
	NeuronA *Neuron
	NeuronB *Neuron
}

func (a *Axon) RandomWeight() {
	a.Weight = rand.Float64()
}
