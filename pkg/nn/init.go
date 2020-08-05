// Initialization
package nn

import (
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/zigenzoog/gonn/pkg"
)

type floatType float32

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// New returns a new neural network instance with the default parameters
func New(reader ...io.Reader) NeuralNetwork {
	n := new(NN)
	if len(reader) > 0 {
		switch r := reader[0].(type) {
		case jsonType:
			if len(r) > 0 {
				//n =
			} else {
				log.Fatal("Отсутствует название файла нейросети")
			}
		case xmlType:
		case csvType:
		case dbType:
		default:
		}
	} else {
		n = &NN{
			Architecture: &perceptron{},
			IsInit:       false,
			IsTrain:      false,
			json:         "",
			xml:          "",
			csv:          "",
		}
		n.Perceptron() //???
	}
	return n
}

func (n *NN) init(lenInput int, lenTarget ...interface{}) bool {
	if a, ok := n.Get().(NeuralNetwork); ok {
		n.IsInit = a.init(lenInput, lenTarget...)
	}
	return n.IsInit
}

func (f floatType) Set(...pkg.Setter) {}
func (f floatType) Get(...pkg.Getter) pkg.GetterSetter {
	return nil
}

// getRand return random number from -0.5 to 0.5
func getRand() (r floatType) {
	r = 0
	for r == 0 {
		r = floatType(rand.Float64() - .5)
	}
	return
}

// getLengthData возвращает длину срезов
func getLengthData(data ...[]float64) []interface{} {
	var tmp []interface{}
	defer func() {
		tmp = nil
	}()
	if len(data) > 0 {
		for _, v := range data {
			tmp = append(tmp, len(v))
		}
	}
	return tmp
}