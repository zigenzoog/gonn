package nn

import (
	"fmt"

	"github.com/zigenzoog/gonn/pkg"
)

type biasBool bool

// Bias
func Bias(bias ...bool) pkg.GetSetter {
	if len(bias) > 0 {
		return biasBool(bias[0])
	}
	return biasBool(false)
}

// Bias
func (n *nn) Bias() bool {
	return n.Architecture.(Parameter).Bias()
}

// Set
func (b biasBool) Set(args ...pkg.Setter) {
	if len(args) > 0 {
		if n, ok := args[0].(*nn); ok && !n.IsInit {
			n.Get().Set(b)
		}
	} else {
		errNN(fmt.Errorf("%w set for bias", ErrEmpty))
	}
}

// Get
func (b biasBool) Get(args ...pkg.Getter) pkg.GetSetter {
	if len(args) > 0 {
		if n, ok := args[0].(Architecture); ok {
			return n.Get().Get(b)
		}
	} else {
		return b
	}
	return nil
}
