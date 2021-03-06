package perceptron

import (
	"fmt"
	"log"

	"github.com/zigenzoog/gonn/pkg"
	"github.com/zigenzoog/gonn/pkg/utils"
)

// WriteConfig writes the configuration and weights to the Filer interface object.
func (nn *NN) WriteConfig(name ...string) (err error) {
	if len(name) > 0 {
		switch d := utils.GetFileType(name[0]).(type) {
		case error:
			err = d
		case utils.Filer:
			err = d.Encode(nn)
		}
	} else if nn.config != nil {
		err = nn.config.Encode(nn)
	} else {
		err = pkg.ErrNoArgs
	}

	if err != nil {
		err = fmt.Errorf("write config: %w", err)
		log.Print(err)
	}
	return
}

// WriteConfig writes weights to the Filer interface object.
func (nn *NN) WriteWeight(name string) (err error) {
	switch d := utils.GetFileType(name).(type) {
	case error:
		err = d
	case utils.Filer:
		err = d.Encode(nn.Weights)
	}

	if err != nil {
		err = fmt.Errorf("write weights: %w", err)
		log.Print(err)
	}
	return
}
