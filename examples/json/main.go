package main

import (
	"path/filepath"

	"github.com/zigenzoog/gonn/pkg/nn"
)

func main() {
	// New returns a new neural network from config.
	n := nn.New(filepath.Join("config", "perceptron.json"))

	// Dataset.
	input := []float64{1, 1}
	target := []float64{0}

	// Training dataset.
	_, _ = n.Train(input, target)

	// Writing weights to a file.
	_ = n.WriteWeight("perceptron_weights.json")
}
