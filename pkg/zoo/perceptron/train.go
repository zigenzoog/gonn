package perceptron

import (
	"fmt"
	"log"
	"math"

	"github.com/zigenzoog/gonn/pkg"
)

// MaxIteration the maximum number of iterations after which training is forcibly terminated.
const MaxIteration = 10e+05

var GetMaxIteration = getMaxIteration

func getMaxIteration() int {
	return MaxIteration
}

// Train training dataset.
func (nn *NN) Train(input []float64, target ...[]float64) (loss float64, count int) {
	var err error
	if len(input) > 0 {
		if len(target) > 0 && len(target[0]) > 0 {
			if !nn.isInit {
				nn.Init(len(input), len(target[0]))
			} else {
				if nn.lenInput != len(input) {
					err = fmt.Errorf("invalid number of elements in the input data")
					goto ERROR
				}
				if nn.lenOutput != len(target[0]) {
					err = fmt.Errorf("invalid number of elements in the target data")
					goto ERROR
				}
			}

			_ = copy(nn.input, input)
			_ = copy(nn.output, target[0])

			for count < GetMaxIteration() {
				count++
				nn.calcNeuron()
				switch loss = nn.calcLoss(); {
				case loss < nn.Limit:
					return
				case math.IsNaN(loss):
					log.Panic("train: not optimal neural network parameters")
				}
				nn.calcMiss()
				nn.updWeight()
			}
			return
		} else {
			err = pkg.ErrNoTarget
		}
	} else {
		err = pkg.ErrNoInput
	}

ERROR:
	log.Printf("train: %v\n", err)
	return -1, 0
}
