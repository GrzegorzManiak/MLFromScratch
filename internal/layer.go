package internal

import "fmt"

type Layer struct {
	Name             string
	Neurons          uint64
	Neuron           Neuron
	GeneratedNeurons []*Neuron
	Axons            []*Axon
	ParentLayer      *Layer
	ChildLayer       *Layer
}

func (l *Layer) GenerateNeurons() {
	var generatedNeurons []*Neuron

	var i uint64
	for i = 0; i < l.Neurons; i++ {
		newNeuron := l.Neuron.Create()
		generatedNeurons = append(generatedNeurons, &newNeuron)
	}

	l.GeneratedNeurons = generatedNeurons
}

func (l1 *Layer) BindTo(l2 *Layer) {
	var axons []*Axon

	for _, l1Neuron := range l1.GeneratedNeurons {
		for _, l2Neuron := range l2.GeneratedNeurons {
			newAxon := Axon{NeuronA: l1Neuron, NeuronB: l2Neuron}
			newAxon.RandomWeight()
			axons = append(axons, &newAxon)
		}
	}

	l2.ParentLayer = l1
	l1.ChildLayer = l2
	l1.Axons = axons
}

func (l *Layer) Visualize() {
	fmt.Printf("Layer: %s\n", l.Name)
	fmt.Printf("Neurons: %d\n", l.Neurons)

	// Print each neuron in the layer
	fmt.Println("Neurons:")
	for i, neuron := range l.GeneratedNeurons {
		fmt.Printf("  Neuron %d: %+v\n", i+1, *neuron)
	}

	// Print each axon (connection) between neurons
	if len(l.Axons) > 0 {
		fmt.Println("Axons:")
		for i, axon := range l.Axons {
			fmt.Printf("  Axon %d: NeuronA ID %d -> NeuronB ID %d (Weight: %.2f)\n",
				i+1, (*axon.NeuronA).GetID(), (*axon.NeuronB).GetID(), axon.Weight)
		}
	} else {
		fmt.Println("No axons in this layer.")
	}

	// Show parent and child layer names if they exist
	if l.ParentLayer != nil {
		fmt.Printf("Parent Layer: %s\n", l.ParentLayer.Name)
	}
	if l.ChildLayer != nil {
		fmt.Printf("Child Layer: %s\n", l.ChildLayer.Name)
	}

	fmt.Println()
}
