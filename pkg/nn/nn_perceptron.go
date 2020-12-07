package nn

import (
	"bytes"
	"fmt"
	"math"
)

const perceptronName = "perceptron"

// Declare conformity with NeuralNetwork interface
var _ NeuralNetwork = (*perceptron)(nil)

// perceptron
type perceptron struct {
	Parameter `json:"-" xml:"-"`

	// Neural network architecture name
	Name string `json:"name" xml:"name"`

	// Array of the number of neurons in each hidden layer
	Hidden HiddenArrUint `json:"hidden" xml:"hidden>layer"`

	// The neuron bias, false or true
	Bias biasBool `json:"bias" xml:"bias"`

	// Activation function mode
	Activation uint8 `json:"activation" xml:"activation"`

	// The mode of calculation of the total error
	Loss uint8 `json:"loss" xml:"loss"`

	// Minimum (sufficient) limit of the average of the error during training
	Limit float64 `json:"limit" xml:"limit"`

	// Learning coefficient, from 0 to 1
	Rate FloatType `json:"rate" xml:"rate"`

	// Weight value
	Weights Float3Type `json:"weights" xml:"weights>weights"`

	// Neuron
	neuron [][]*perceptronNeuron

	lastLayerIndex int
	lenInput       int
	lenOutput      int

	// State of the neural network
	isInit  bool // Neural network initializing flag
	isTrain bool // Neural network training flag

	// Config
	jsonName string
}

type perceptronNeuron struct {
	value FloatType
	miss  FloatType
}

// Perceptron return perceptron neural network
func Perceptron() *perceptron {
	return &perceptron{
		Name:       perceptronName,
		Hidden:     HiddenArrUint{},
		Bias:       false,
		Activation: ModeSIGMOID,
		Loss:       ModeMSE,
		Limit:      .01,
		Rate:       FloatType(DefaultRate),
	}
}

func (p *perceptron) name() string {
	return p.Name
}

func (p *perceptron) setName(name string) {
	p.Name = name
}

func (p *perceptron) stateInit() bool {
	return p.isInit
}

func (p *perceptron) setStateInit(state bool) {
	p.isInit = state
}

func (p *perceptron) stateTrain() bool {
	return p.isTrain
}

func (p *perceptron) setStateTrain(state bool) {
	p.isTrain = state
}

func (p *perceptron) nameJSON() string {
	return p.jsonName
}

func (p *perceptron) setNameJSON(name string) {
	p.jsonName = name
}

// HiddenLayer
func (p *perceptron) HiddenLayer() []uint {
	return p.Hidden
}

// SetHiddenLayer
func (p *perceptron) SetHiddenLayer(layer []uint) {
	p.Hidden = layer
}

// NeuronBias
func (p *perceptron) NeuronBias() bool {
	return bool(p.Bias)
}

// SetNeuronBias
func (p *perceptron) SetNeuronBias(bias bool) {
	p.Bias = biasBool(bias)
}

// ActivationMode
func (p *perceptron) ActivationMode() uint8 {
	return p.Activation
}

// SetActivationMode
func (p *perceptron) SetActivationMode(mode uint8) {
	p.Activation = mode
}

// LossMode
func (p *perceptron) LossMode() uint8 {
	return p.Loss
}

// SetLossMode
func (p *perceptron) SetLossMode(mode uint8) {
	p.Loss = mode
}

// LossLimit
func (p *perceptron) LossLimit() float64 {
	return p.Limit
}

// SetLossLimit
func (p *perceptron) SetLossLimit(limit float64) {
	p.Limit = limit
}

// LearningRate
func (p *perceptron) LearningRate() float32 {
	return float32(p.Rate)
}

// SetLearningRate
func (p *perceptron) SetLearningRate(rate float32) {
	p.Rate = FloatType(rate)
}

// Weight
func (p *perceptron) Weight() Floater {
	return &p.Weights
}

// SetWeight
func (p *perceptron) SetWeight(weight Floater) {
	p.Weights = weight.(Float3Type)
}

// Set
func (p *perceptron) Set(args ...Setter) {
	if len(args) > 0 {
		for _, a := range args {
			switch v := a.(type) {
			case HiddenArrUint:
				p.Hidden = v
			case biasBool:
				p.Bias = v
			case activationModeUint:
				p.Activation = uint8(v)
			case lossModeUint:
				p.Loss = uint8(v)
			case lossLimitFloat:
				p.Limit = float64(v)
			case rateFloat:
				p.Rate = FloatType(v)
			//case *weight:
			//p.setWeight(v.buffer.(*Float3Type))
			default:
				LogError(fmt.Errorf("%T %w for perceptron", v, ErrMissingType))
			}
		}
	} else {
		LogError(fmt.Errorf("%w set for perceptron", ErrEmpty))
	}
}

// Get
func (p *perceptron) Get(args ...Getter) GetSetter {
	if len(args) > 0 {
		for _, a := range args {
			switch a.(type) {
			case HiddenArrUint:
				return p.Hidden
			case biasBool:
				return p.Bias
			case activationModeUint:
				return activationModeUint(p.Activation)
			case lossModeUint:
				return lossModeUint(p.Loss)
			case lossLimitFloat:
				return lossLimitFloat(p.Limit)
			case rateFloat:
				return p.Rate
			//case *weight:
			//return p.getWeight()
			default:
				LogError(fmt.Errorf("%T %w for perceptron", a, ErrMissingType))
			}
		}
	}
	return p
}

// Read
func (p *perceptron) Read(reader Reader) {
	switch r := reader.(type) {
	case Filer:
		r.Read(p)
	default:
		LogError(fmt.Errorf("%T %w for read: %v", r, ErrMissingType, r))
	}
}

// Write
func (p *perceptron) Write(writer ...Writer) {
	if len(writer) > 0 {
		for _, w := range writer {
			switch v := w.(type) {
			case Filer:
				v.Write(p)
			case *report:
				p.writeReport(v)
			default:
				LogError(fmt.Errorf("%T %w for write: %v", v, ErrMissingType, w))
			}
		}
	} else {
		LogError(fmt.Errorf("%w for write", ErrEmpty))
	}
}

// init initialize
func (p *perceptron) init(lenInput, lenTarget int) bool {
	p.lastLayerIndex = len(p.Hidden)
	p.lenInput = lenInput
	p.lenOutput = lenTarget
	layers := append(p.Hidden, uint(p.lenOutput))
	lenLayer := len(layers)

	bias := 0
	if p.Bias {
		bias = 1
	}
	biasInput := p.lenInput + bias
	var biasLayer int

	p.Weights = make(Float3Type, lenLayer)
	p.neuron = make([][]*perceptronNeuron, lenLayer)

	for i, l := range layers {
		p.Weights[i] = make([][]FloatType, l)
		p.neuron[i] = make([]*perceptronNeuron, l)
		if i > 0 {
			biasLayer = int(layers[i-1]) + bias
		}
		for j := 0; j < int(l); j++ {
			if i > 0 {
				p.Weights[i][j] = make([]FloatType, biasLayer)
			} else {
				p.Weights[i][j] = make([]FloatType, biasInput)
			}
			for k := range p.Weights[i][j] {
				p.Weights[i][j][k] = getRand()
			}
			p.neuron[i][j] = &perceptronNeuron{}
		}
	}
	return true
}

// reInit
func (p *perceptron) reInit() {
	bias := 0
	if p.Bias {
		bias = 1
	}
	length := len(p.Weights) - 1
	p.Hidden = make(HiddenArrUint, length)
	for i := range p.Hidden {
		p.Hidden[i] = uint(len(p.Weights[i]))
	}
	p.isInit = p.init(len(p.Weights[0][0])-bias, len(p.Weights[length]))
}

// calcNeuron
func (p *perceptron) calcNeuron(input []float64) {
	wait := make(chan bool)
	defer close(wait)

	var length int
	for i, v := range p.neuron {
		fmt.Println(i, v)
		dec := i - 1
		if i > 0 {
			length = len(p.neuron[dec])
		} else {
			length = p.lenInput
		}
		for j, n := range v {
			go func(j int, n *perceptronNeuron) {
				n.value = 0
				for k, w := range p.Weights[i][j] {
					if k < length {
						if i > 0 {
							n.value += p.neuron[dec][k].value * w
						} else {
							n.value += FloatType(input[k]) * w
						}
					} else {
						n.value += w
					}
				}
				n.value = FloatType(calcActivation(float64(n.value), p.Activation))
				wait <- true
			}(j, n)
		}
		for range v {
			<-wait
		}
	}
}

// calcLoss calculating the error of the output neuron
func (p *perceptron) calcLoss(target []float64) (loss float64) {
	for i, n := range p.neuron[p.lastLayerIndex] {
		n.miss = FloatType(target[i]) - n.value
		switch p.Loss {
		default:
			fallthrough
		case ModeMSE, ModeRMSE:
			loss += math.Pow(float64(n.miss), 2)
		case ModeARCTAN:
			loss += math.Pow(math.Atan(float64(n.miss)), 2)
		}
		n.miss *= FloatType(calcDerivative(float64(n.miss), p.Activation))
	}
	loss /= float64(p.lenOutput)
	if p.Loss == ModeRMSE {
		loss = math.Sqrt(loss)
	}
	return
}

// calcMiss calculating the error of neurons in hidden layers
func (p *perceptron) calcMiss() {
	wait := make(chan bool)
	defer close(wait)

	for i := p.lastLayerIndex - 1; i >= 0; i-- {
		inc := i + 1
		for j, n := range p.neuron[i] {
			go func(j int, n *perceptronNeuron) {
				n.miss = 0
				for k, m := range p.neuron[inc] {
					n.miss += m.miss * p.Weights[inc][k][j]
				}
				n.miss *= FloatType(calcDerivative(float64(n.value), p.Activation))
				wait <- true
			}(j, n)
		}
		for range p.neuron[i] {
			<-wait
		}
	}
}

// updWeight update weights
func (p *perceptron) updWeight(input []float64) {
	wait := make(chan bool)
	defer close(wait)

	var length int
	for i, u := range p.Weights {
		dec := i - 1
		if i > 0 {
			length = len(p.neuron[dec])
		} else {
			length = p.lenInput
		}
		for j, v := range u {
			go func(i, j, dec, length int, grad FloatType, v []FloatType) {
				for k := range v {
					if k < length {
						if i > 0 {
							p.Weights[i][j][k] += p.neuron[dec][k].value * grad
						} else {
							p.Weights[i][j][k] += FloatType(input[k]) * grad
						}
					} else {
						p.Weights[i][j][k] += grad
					}
				}
				wait <- true
			}(i, j, dec, length, p.neuron[i][j].miss*p.Rate, v)
		}
	}
	for _, u := range p.Weights {
		for range u {
			<-wait
		}
	}
}

// Train training neural network
func (p *perceptron) Train(input []float64, target ...[]float64) (loss float64, count int) {
	if !p.isInit {
		if p.isInit = p.init(len(input), len(target[0])); !p.isInit {
			LogError(fmt.Errorf("%w for train", ErrInit))
			return -1, 0
		}
	}
	fmt.Println(p)
	//_ = copy(p.input, input)
	//_ = copy(p.target, target[0])
	if len(target) > 0 {
		for count < MaxIteration {
			p.calcNeuron(input)
			if loss = p.calcLoss(target[0]); loss <= p.Limit {
				break
			}
			p.calcMiss()
			p.updWeight(input)
			count++
		}
	} else {
		LogError(ErrNoTarget)
		return -1, 0
	}
	if !p.isTrain && count > 0 {
		p.isTrain = true
	}
	return
}

// Verify verifying neural network
func (p *perceptron) Verify(input []float64, target ...[]float64) (loss float64) {
	if !p.isInit {
		if p.isInit = p.init(len(input), len(target[0])); !p.isInit {
			LogError(fmt.Errorf("%w for verify", ErrInit))
			return -1
		}
	}
	if len(target) > 0 {
		p.calcNeuron(input)
		loss = p.calcLoss(target[0])
	} else {
		LogError(ErrNoTarget)
		return -1
	}
	return
}

// Query querying neural network
func (p *perceptron) Query(input []float64) (output []float64) {
	if !p.isTrain {
		LogError(fmt.Errorf("query: %w", ErrNotTrained))
		if !p.isInit {
			LogError(fmt.Errorf("%w for query", ErrInit))
		}
		return nil
	}
	p.calcNeuron(input)
	output = make([]float64, p.lenOutput)
	for i, n := range p.neuron[p.lastLayerIndex] {
		output[i] = float64(n.value)
	}
	return
}

// initWeight
/*func (p *perceptron) initWeight() {
	p.Weights = make(Float3Type, len(p.axon))
	for i, v := range p.axon {
		p.Weights[i] = make(Float2Type, len(p.axon[i]))
		for j := range v {
			p.Weights[i][j] = make(Float1Type, len(p.axon[i][j]))
		}
	}
	p.weight = &weight{
		isInitWeight: true,
		buffer:       &p.Weights,
	}
}*/

// getWeight
/*func (p *perceptron) getWeight() *Float3Type {
	p.copyWeight()
	return &p.Weights
}*/

// setWeight
/*func (p *perceptron) setWeight(weight *Float3Type) {
	for i, u := range *weight {
		for j, v := range u {
			for k, w := range v {
				p.axon[i][j][k].weight = w
			}
		}
	}
}*/

// copyWeight copies weights to the buffer
/*func (p *perceptron) copyWeight() {
	if p.weight == nil {
		p.initWeight()
	}
	for i, u := range p.axon {
		for j, v := range u {
			for k, w := range v {
				p.Weights[i][j][k] = w.weight
			}
		}
	}
}*/

// pasteWeight inserts weights from the buffer
/*func (p *perceptron) pasteWeight() (err error) {
	if p.Weights != nil {
		for i, u := range p.Weights {
			for j, v := range u {
				for k, w := range v {
					p.axon[i][j][k].weight = w
				}
			}
		}
		p.deleteWeight()
	} else {
		err = fmt.Errorf("paste weight error: missing weights")
	}
	return
}*/

// deleteWeight
/*func (p *perceptron) deleteWeight() {
	p.Weights = nil
	p.weight = nil
}*/

// readJSON
/*func (p *perceptron) readJSON(value interface{}) {
	if b, err := json.Marshal(&value); err != nil {
		LogError(fmt.Errorf("read marshal %w", err))
	} else if err = json.Unmarshal(b, &p); err != nil {
		LogError(fmt.Errorf("read unmarshal %w", err))
	}
	p.reInit()
	if err := p.pasteWeight(); err != nil {
		LogError(fmt.Errorf("read json: %w", err))
	}
}*/

// writeReport report of neural network training results in io.Writer
func (p *perceptron) writeReport(rep *report) {
	s := "----------------------------------------------\n"
	n := "\n"
	m := "\n\n"
	b := bytes.NewBufferString("Report of Perceptron Neural Network\n\n")

	printFormat := func(format string, a ...interface{}) {
		if _, err := fmt.Fprintf(b, format, a...); err != nil {
			LogError(fmt.Errorf("write report error: %w", err))
		}
	}

	// Input layer
	if in, ok := rep.args[0].([]float64); ok {
		printFormat("%s0 Input layer size: %d\n%sNeurons:\t", s, p.lenInput, s)
		for _, v := range in {
			printFormat("  %v", v)
		}
		printFormat("%s", m)
	}

	// Layers: neuron, miss
	var t string
	for i, v := range p.neuron {
		switch i {
		case p.lastLayerIndex:
			t = "Output layer"
		default:
			t = "Hidden layer"
		}
		printFormat("%s%d %s size: %d\n%sNeurons:\t", s, i+1, t, len(p.neuron[i]), s)
		for _, w := range v {
			printFormat("  %11.8f", w.value)
		}
		printFormat("\nMiss:\t\t")
		for _, w := range v {
			printFormat("  %11.8f", w.miss)
		}
		printFormat("%s", m)
	}

	// Axons: weight
	printFormat("%sAxons (weights)\n%s", s, s)
	for _, u := range p.Weights {
		for i, v := range u {
			printFormat("%d", i+1)
			for _, w := range v {
				printFormat("\t%11.8f", w)
			}
			printFormat("%s", n)
		}
		printFormat("%s", n)
	}

	// Resume
	if loss, ok := rep.args[1].(float64); ok {
		printFormat("%sTotal loss (error):\t\t%v\n", s, loss)
	}
	if count, ok := rep.args[2].(int); ok {
		printFormat("Number of iteration:\t%v\n", count)
	}

	if _, err := b.WriteTo(rep.file); err != nil {
		LogError(fmt.Errorf("write report error: %w", err))
	} else if err = rep.file.Close(); err != nil {
		LogError(fmt.Errorf("write report close error: %w", err))
	}
}
