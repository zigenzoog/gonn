package nn

import (
	"reflect"
	"testing"

	"github.com/zigenzoog/gonn/pkg"
	"github.com/zigenzoog/gonn/pkg/zoo"
)

func TestNew(t *testing.T) {
	testNN := zoo.Get(zoo.Perceptron)
	tests := []struct {
		name string
		gave []string
		want pkg.NeuralNetwork
	}{
		{
			name: "#1_warning_empty",
			gave: []string{},
			want: testNN,
		},
		{
			name: "#2_" + zoo.Perceptron,
			gave: []string{zoo.Perceptron},
			want: testNN,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.gave...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New()\ngot:\t%v\nwant:\t%v", got, tt.want)
			}
		})
	}
}
