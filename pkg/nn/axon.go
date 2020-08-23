package nn

import (
	"github.com/zigenzoog/gonn/pkg"
)

type weightType floatType

//
type axon struct {
	weight  floatType             //
	synapse map[string]pkg.Getter //
}

func (a axon) Set(...pkg.Setter) {}
func (a axon) Get(...pkg.Getter) pkg.GetterSetter {
	return a
}

// Weight
func Weight() pkg.GetterSetter {
	return weightType(0)
}

func (w weightType) Set(args ...pkg.Setter) {
	if len(args) > 0 {
		if n, ok := args[0].(NeuralNetwork); ok {
			n.Get().Set(w)
		}
	} else {
		pkg.Log("Empty set", true) // !!!
	}
}

func (w weightType) Get(args ...pkg.Getter) pkg.GetterSetter {
	if len(args) > 0 {
		if a, ok := args[0].(NeuralNetwork); ok {
			return a.Get().Get(w)
		}
	} else {
		return w
	}
	return nil
}

func (w weightType) Copy(copier pkg.Getter) {
	if n, ok := copier.(*NN); ok {
		if a, ok := n.Architecture.(NeuralNetwork); ok {
			a.Copy(w)
		}
	}
}

func (w weightType) Paste(paster pkg.Getter) (err error) {
	if n, ok := paster.(*NN); ok {
		if a, ok := n.Architecture.(NeuralNetwork); ok {
			err = a.Paste(w)
		}
	}
	return
}

func getSynapseInput(axon *axon) (input floatType) {
	switch s := axon.synapse["input"].(type) {
	case floatType:
		input = s
	case biasType:
		if s {
			input = 1
		}
	case *neuron:
		input = s.value
	default:
		panic("error!!!") // !!!
	}
	return
}

/*func Axon() GetterSetter {
	return &axon{}
}

// Set
func (a *axon) Set(args ...Setter) {
	if len(args) > 0 {
		if n, ok := args[0].(NeuralNetwork); ok {
			n.Get().Set(a)
		}
	} else {
		Log("Empty Set()", true) // !!!
	}
}

// Get
func (a *axon) Get(args ...Getter) GetterSetter {
	if len(args) > 0 {
		if n, ok := args[0].(NeuralNetwork); ok {
			return n.Get().Get(a)
		}
	} else {
		return a
	}
	return nil
}*/