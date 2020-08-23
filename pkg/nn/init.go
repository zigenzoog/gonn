// Initialization
package nn

import (
	"math/rand"
	"time"

	"github.com/zigenzoog/gonn/pkg"
)

type floatType float32

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// New returns a new neural network instance with the default parameters
func New(reader ...pkg.Reader) NeuralNetwork {
	n := &NN{
		IsInit:  false,
		IsTrain: false,
		json:    "",
		xml:     "",
		csv:     "",
		db:      "",
	}
	if len(reader) > 0 {
		switch r := reader[0].(type) {
		case *perceptron:
			n.Architecture = r
			n.perceptron()
		/*case jsonType:
			if len(r) > 0 {
				//n =
			} else {
				log.Fatal("Отсутствует название файла нейросети")
			}
		case xmlType:
		case csvType:
		case dbType:*/
		default:
		}
	} else {
		n.Architecture = &perceptron{}
		n.perceptron()
	}
	return n
}

func Reset() *NN {
	n := &NN{
		Architecture: nil,
		IsInit:       false,
		IsTrain:      false,
		json:         "",
		xml:          "",
		csv:          "",
		db:           "",
	}
	return n
}

func (n *NN) erase() {
	n.Architecture = nil
	n.IsInit = false
	n.IsTrain = false
	n.json = ""
	n.xml = ""
	n.csv = ""
	n.db = ""
}

func (n *NN) init(lenInput int, lenTarget ...interface{}) bool {
	if a, ok := n.Get().(NeuralNetwork); ok {
		n.IsInit = a.init(lenInput, lenTarget...)
	}
	return n.IsInit
}

func (f floatType) Set(...pkg.Setter) {}
func (f floatType) Get(...pkg.Getter) pkg.GetterSetter {
	return f
}

// getRand return random number from -0.5 to 0.5
func getRand() (r floatType) {
	r = 0
	for r == 0 {
		r = floatType(rand.Float64() - .5)
	}
	return
}

// getLengthData returns the length of the slices
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
	//fmt.Println(tmp)
	return tmp
}

func Debug(args ...interface{}) {
	//fmt.Println(getLengthData([]float64{.2,.3,.8}))
}