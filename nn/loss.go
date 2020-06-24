// Loss
package nn

type (
	modeLossType uint8		//
	lossType     floatType	// Level loss
)

const (
	ModeMSE      modeLossType = 0		// Mean Squared Error
	ModeRMSE     modeLossType = 1		// Root Mean Squared Error
	ModeARCTAN   modeLossType = 2		// Arctan
	MinLevelLoss lossType     = 10e-33	// The minimum value of the error limit at which training is forcibly terminated
)

func Loss(mode ...modeLossType) Setter {
	if len(mode) == 0 {
		return modeLossType(0)
	} else {
		return mode[0]
	}
}

func LevelLoss(level ...lossType) Setter {
	if len(level) == 0 {
		return lossType(0)
	} else {
		return level[0]
	}
}

// Setter
func (m modeLossType) Set(set ...Setter) {
	if v, ok := getArchitecture(set[0]); ok {
		if c, ok := m.Check().(modeLossType); ok {
			v.Set(c)
		}
	}
}

func (l lossType) Set(set ...Setter) {
	if v, ok := getArchitecture(set[0]); ok {
		if c, ok := l.Check().(lossType); ok {
			v.Set(c)
		}
	}
}

// Getter
func (m modeLossType) Get(set ...Setter) Getter {
	if v, ok := getArchitecture(set[0]); ok {
		return v.Get(m)
	}
	return nil
}

func (l lossType) Get(set ...Setter) Getter {
	if v, ok := getArchitecture(set[0]); ok {
		return v.Get(l)
	}
	return nil
}

// Checker
func (m modeLossType) Check() Getter {
	switch {
	case m < 0 || m > ModeARCTAN:
		return ModeMSE
	default:
		return m
	}
}

func (l lossType) Check() Getter {
	switch {
	case l < 0:
		return MinLevelLoss
	default:
		return l
	}
}