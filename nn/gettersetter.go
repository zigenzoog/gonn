package nn

import (
	"math/rand"
)

type GetterSetter interface {
	Getter
	Setter
}

type Getter interface {
	Get(...Setter) Getter
}

type Setter interface {
	Set(...Setter)
}

// Setter
func (n *nn) Set(args ...Setter) {
	if len(args) == 0 {
		Log("Empty Set()", false)
	} else {
		for _, v := range args {
			if s, ok := v.(Setter); ok {
				s.Set(n)
			}
		}
	}
}

// Getter
func (n *nn) Get(args ...Setter) Getter {
	if len(args) == 0 {
		return n
	} else {
		for _, v := range args {
			if g, ok := v.(Getter); ok {
				return g.Get(n)
			}
		}
	}
	return nil
}

//
/*func  (i intType) Set(set ...Setter) {
	Log("", false)
}*/

func  (f floatType) Set(args ...Setter) {
	Log("", false) // !!!
}

func  (f floatType) Get(args ...Setter) Getter {
	Log("", false) // !!!
	return nil
}

func  (f FloatType) Set(args ...Setter) {
	Log("", false) // !!!
}

func  (f FloatType) Get(args ...Setter) Getter {
	Log("", false) // !!!
	return nil
}

//
func getArchitecture(net Setter) (NeuralNetwork, bool) {
	if n, ok := net.(*nn); ok {
		if v, ok := n.architecture.(NeuralNetwork); ok {
			return v, ok
		}
	}
	return nil, false
}

// Return random number from -0.5 to 0.5
func getRand() (r floatType) {
	r = 0
	for r == 0 {
		r = floatType(rand.Float64() - .5)
	}
	return
}