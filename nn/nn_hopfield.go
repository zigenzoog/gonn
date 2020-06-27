// Hopfield Neural Network - under construction
package nn

/*type Hopfield interface {
	Hopfield() NeuralNetwork
}*/

type hopfield struct {
	Architecture
	Processor
}

// Initializing Hopfield Neural Network
func (n *nn) Hopfield() NeuralNetwork {
	n.architecture = &hopfield{
		Architecture:	n,
	}
	return n
}

// Setter
func (h *hopfield) Set(set ...Setter) {
	switch /*v :=*/ set[0].(type) {
	/*case BiasType:
		r.bias = v
	case RateType:
		r.rate = v
	case HiddenType:
		r.hiddenLayer = v*/
	default:
		Log("This type of variable is missing for Hopfield Neural Network", false)
	}
}

// Getter
func (h *hopfield) Get(set ...Setter) Getter {
	switch set[0].(type) {
	/*case Bias:
		return r.bias
	case Rate:
		return r.rate
	case Hidden:
		return r.hiddenLayer*/
	default:
		Log("This type of variable is missing for Perceptron Neural Network", false)
		return nil
	}
}

// Train
/*func (h *hopfield) Train(input, target []float64) (loss float64, count int) {
	return
}

// Query
func (h *hopfield) Query(input []float64) []float64 {
	panic("implement me")
}*/