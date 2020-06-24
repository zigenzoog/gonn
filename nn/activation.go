//
package nn

import "math"

type modeActivationType uint8

const (
	ModeLINEAR    modeActivationType = 0 // Linear/identity
	ModeRELU      modeActivationType = 1 // ReLu - rectified linear unit
	ModeLEAKYRELU modeActivationType = 2 // Leaky ReLu - leaky rectified linear unit
	ModeSIGMOID   modeActivationType = 3 // Logistic, a.k.a. sigmoid or soft step
	ModeTANH      modeActivationType = 4 // TanH - hyperbolic
)

func Activation(mode ...modeActivationType) Setter {
	if len(mode) == 0 {
		return modeActivationType(0)
	} else {
		return mode[0]
	}
}

// Setter
func (m modeActivationType) Set(set ...Setter) {
	if v, ok := getArchitecture(set[0]); ok {
		if c, ok := m.Check().(modeActivationType); ok {
			v.Set(c)
		}
	}
}

// Getter
func (m modeActivationType) Get(set ...Setter) Getter {
	if v, ok := getArchitecture(set[0]); ok {
		return v.Get(m)
	}
	return nil
}

// Checker
func (m modeActivationType) Check() Getter {
	switch {
	case m < 0 || m > ModeTANH:
		return ModeSIGMOID
	default:
		return m
	}
}

type Activator interface {
	Get(float64) float64
}

type activation struct {
	Mode modeActivationType
}

type derivative struct {
	Mode modeActivationType
}

func getActivation(value float64, a Activator) float64 {
	return a.Get(value)
}

// Activation function
func (a *activation) Get(value float64) float64 {
	switch a.Mode {
	default:
		fallthrough
	case ModeLINEAR:
		return value
	case ModeRELU:
		switch {
		case value < 0:
			return 0
		default:
			return value
		}
	case ModeLEAKYRELU:
		switch {
		case value < 0:
			return .01 * value
		default:
			return value
		}
	case ModeSIGMOID:
		return 1 / (1 + math.Exp(-value))
	case ModeTANH:
		value = math.Exp(2 * value)
		if math.IsInf(value, 1) {
			return 1
		}
		return (value - 1) / (value + 1)
	}
}

// Derivative activation function
func (d *derivative) Get(value float64) float64 {
	switch d.Mode {
	default:
		fallthrough
	case ModeLINEAR:
		return 1
	case ModeRELU:
		switch {
		case value < 0:
			return 0
		default:
			return 1
		}
	case ModeLEAKYRELU:
		switch {
		case value < 0:
			return .01
		default:
			return 1
		}
	case ModeSIGMOID:
		return value * (1 - value)
	case ModeTANH:
		return 1 - math.Pow(value, 2)
	}
}
