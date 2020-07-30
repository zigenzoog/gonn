//
package nn

import "github.com/zigenzoog/gonn/pkg"

// Declare conformity with NeuralNetwork interface
var _ NeuralNetwork = (*NN)(nil)

//
type NeuralNetwork interface {
	//
	Architecture

	//
	Constructor

	//
	pkg.GetterSetter
}

//
type NN struct {
	Architecture // Architecture of neural network

	isInit  bool // Neural network initializing flag
	isTrain bool // Neural network training flag

	json	string
	xml		string
	/*csv		string*/
}