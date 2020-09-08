package nn

import (
	"fmt"
)

// The maximum number of iterations after which training is forcibly terminated
const MaxIteration int = 10e+05

// Train
func (n *nn) Train(input []float64, target ...[]float64) (loss float64, count int) {
	if !n.IsInit {
		if n.IsInit = n.init(len(input), getLengthData(target...)...); !n.IsInit {
			errNN(fmt.Errorf("%w for train", ErrInit))
			return
		}
	}
	if a, ok := n.Architecture.(NeuralNetwork); ok {
		loss, count = a.Train(input, target...)
		if count > 0 {
			n.IsTrain = true
		}
	}
	return
}