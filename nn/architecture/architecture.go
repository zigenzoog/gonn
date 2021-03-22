package architecture

import (
	"fmt"
	"log"
	"strings"

	"github.com/teratron/gonn"
	"github.com/teratron/gonn/nn/hopfield"
	"github.com/teratron/gonn/nn/perceptron"
)

const (
	Perceptron = perceptron.Name
	Hopfield   = hopfield.Name
)

// Get
func Get(name ...string) gonn.NeuralNetwork {
	if len(name) > 0 {
		switch strings.ToLower(name[0]) {
		case Perceptron:
			return perceptron.Perceptron()
		case Hopfield:
			return hopfield.Hopfield()
		default:
			log.Fatal(fmt.Errorf("get architecture: neural network is %w", gonn.ErrNotRecognized))
			//return nil
		}
	}
	return perceptron.Perceptron()
}