package main

import (
	"fmt"

	"github.com/zigenzoog/gonn/pkg/nn"
)

/*type Parent struct {
	child Child
}
type Child struct {
	num int
}

func (p Parent) print() {
	fmt.Println(p.child.num)
}
func (c Child) print() {
	fmt.Println(c.num + 1)
}*/

//type Layer []Layer

func main() {
	/*parent := Parent{Child{2}}
	defer parent.print()*/

	/*var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)
	i = 0.42
	fmt.Printf("(%v, %T)\n", i, i)
	b := i.(float64)
	fmt.Printf("(%v, %T)\n", b, b)*/

	/*var mx nn.Matrix
	mx.Bias = 0.1
	fmt.Println(mx.Bias)
	a := nn.Activation{Mode: nn.TANH}
	fmt.Println(a, a.Get(.2))

	neu := nn.Neuron{
		X:	1,
		Y:	0,
		Value:	0.1,
		N:	&nn.PrevNeuronLayer{1, .2, .3},
		W:	&nn.PrevWeightLayer{.3, 52},
	}
	fmt.Println(neu, *neu.N, *neu.W)*/

	/*mx.Neuron  = make([][]nn.Neuron, 1)
	mx.Neuron[0] = make([]nn.Neuron, 1)
	mx.Neuron[0][0].Value = 0.02
	fmt.Println(mx.Neuron[0][0].Get())*/

	/*rgba := color.RGBA{0,0,0,255}
	fmt.Println(rgba)

	clr := color.Color(rgba)
	fmt.Println(clr)*/

	/*var mx nn.Matrix
	mx.Axon = make([]nn.Axon, 1)
	mx.Axon[0].Synapse = make(map[string]nn.Neuroner)

	a := nn.Bias(4)
	fmt.Printf("%T %v\n", a, a)
	mx.Axon[0].Synapse["bias"] = &a

	if b, ok := mx.Axon[0].Synapse["bias"]; ok && *b.(*nn.Bias) == 4 {
		fmt.Printf("%T, %v\n", b, *b.(*nn.Bias))
	}

	n := nn.Neuron{Value: .5}
	fmt.Printf("%T %v\n", n, n)

	mx.Axon[0].Synapse["input"] = n
	fmt.Printf("%T %v\n", mx.Axon[0].Synapse["input"], mx.Axon[0].Synapse["input"])

	c := nn.Axon{
		Synapse: map[string]nn.Neuroner{
			"input": n,
			"bias": &a,
		},
	}
	fmt.Printf("%T %v\n", c, c)

	in := c.Synapse["input"]
	fmt.Printf("%T %v\n", in, in.(*nn.Neuron).Value) // method
	//fmt.Printf("%T %v\n", in, in.(nn.Neuron).Value) // struct {Neuroner}*/

	/*var mx nn.Matrix
	fmt.Println(mx.IsInit)*/

	// Application
	a := nn.App()
	fmt.Println("nn.App():", a)

	// Neural Network
	n := nn.New()	// same n := nn.New().Perceptron()
	//n := nn.New().Perceptron()
	//n := nn.New().RadialBasis()
	//n := nn.New().Hopfield()
	fmt.Println("nn.New():", n)

	// Common
	n.Set(
		nn.Bias(false),
		nn.Rate(.33),
		nn.ModeActivation(nn.ModeTANH),
		nn.ModeLoss(nn.ModeARCTAN),
		nn.LevelLoss(.0005),
		nn.HiddenLayer(1, 5, 9),

		/*nn.Language("ru"),
		nn.Logging(1)*/) //set
	fmt.Printf("n.Get(): %T %v\n", n.Get(), n.Get()) //get

	// Bias
	n.Set(nn.Bias(true)) //set
	fmt.Println("n.Get(nn.Bias()):", n.Get(nn.Bias())) //get

	// Rate
	n.Set(nn.Rate(.1)) //set
	fmt.Println("n.Get(nn.Rate()):", n.Get(nn.Rate())) //get

	// Activation
	n.Set(nn.ModeActivation(nn.ModeSIGMOID)) //set
	fmt.Println("n.Get(nn.ModeActivation()):", n.Get(nn.ModeActivation())) //get

	// Loss
	n.Set(nn.ModeLoss(nn.ModeMSE)) //set
	fmt.Println("n.Get(nn.ModeLoss()):", n.Get(nn.ModeLoss())) //get

	// Level loss
	n.Set(nn.LevelLoss(.04)) //set
	fmt.Println("n.Get(nn.LevelLoss()):", n.Get(nn.LevelLoss())) //get

	// Hidden layers
	n.Set(nn.HiddenLayer(3, 2)) //set
	fmt.Println("n.Get(nn.HiddenLayer()):", n.Get(nn.HiddenLayer())) //get

	// Language
	//n.Set(nn.Language("ru")) //set
	//fmt.Println("n.Get(nn.Language()):", n.Get(nn.Language())) //get

	// Logging
	//n.Set(nn.Logging(0)) //set
	//fmt.Println("n.Get(nn.Logging()):", n.Get(nn.Logging())) //get

	//
	input  := []float64{2.3, 3.1}
	target := []float64{3.6}

	//
	fmt.Println(n.Train(input, target))

	//
	//fmt.Println(n.Query(input))

	//
	//fmt.Println(n.Verify(input, target))

	//fmt.Println(n.Get(nn.Neuron()))
	//fmt.Printf("++++ Act: %.4f\n", 100*calcActivation(1, ModeSIGMOID))
}