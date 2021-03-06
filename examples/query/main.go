package main

import (
	"fmt"
	"path/filepath"

	"github.com/zigenzoog/gonn/pkg/nn"
)

func main() {
	// New returns a new neural network from config.
	n := nn.New(filepath.Join("config", "perceptron.json"))

	// Input dataset.
	input := []float64{.27, .31, .52}

	// Getting the results of the trained network.
	output := n.Query(input)
	fmt.Println(output)
}
