//
package nn

import "github.com/zigenzoog/gonn/pkg"

//type neuronType [][]*neuron

type Neuron struct {
	value    floatType // Neuron value
	axon     []*Axon
	specific pkg.Getter
}

/*type miss struct {
	miss floatType // Error value
}*/

/*func Neuron() GetterSetter {
	return &neuron{}
}*/

// Setter
func (n *Neuron) Set(args ...pkg.Setter) {
	/*if len(args) > 0 {
		if a, ok := args[0].(NeuralNetwork); ok {
			a.Get().Set(n)
		}
	} else {
		Log("Empty Set()", true) // !!!
	}*/
}

// Getter
func (n *Neuron) Get(args ...pkg.Getter) pkg.GetterSetter {
	if len(args) > 0 {
		if a, ok := args[0].(NeuralNetwork); ok {
			return a.Get().Get(n)
		}
	} else {
		return n
	}
	return nil
}

/*func (m *miss) Get(...Getter) GetterSetter {
	return nil
}*/

/*func (n *neuronType) Set(args ...Setter) {
}

func (n *neuronType) Get(args ...Getter) GetterSetter {
	return nil
}*/