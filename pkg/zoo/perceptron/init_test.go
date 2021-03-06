package perceptron

import (
	"reflect"
	"testing"

	"github.com/zigenzoog/gonn/pkg"
	"github.com/zigenzoog/gonn/pkg/params"
)

func TestNN_Init(t *testing.T) {
	tests := []struct {
		name string
		gave []interface{}
		got  *NN
		want *NN
	}{
		{
			name: "#1_error_type",
			gave: []interface{}{"test_error"},
			got:  &NN{},
			want: &NN{},
		},
		{
			name: "#2_empty_arguments",
			gave: []interface{}{},
			got:  &NN{},
			want: &NN{},
		},
	}
	for _, tt := range tests {
		tt.got.Weights = tt.want.Weights
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.Init(tt.gave...); !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("Init()\ngot:\t%v\nwant:\t%v", tt.got, tt.want)
			}
		})
	}
}

func TestNN_initFromNew(t *testing.T) {
	r := params.GetRandFloat()
	tests := []struct {
		name string
		got  *NN
		want *NN
	}{
		{
			name: "#1",
			got: &NN{
				Activation: params.ModeSIGMOID,
			},
			want: &NN{
				Bias:       false,
				Hidden:     []int{0},
				Activation: params.ModeSIGMOID,
				Weights: pkg.Float3Type{
					{
						{r, r},
						{r, r},
					},
				},
				neuron: [][]*neuron{
					{
						&neuron{},
						&neuron{},
					},
				},
				lenInput:       2,
				lenOutput:      2,
				lastLayerIndex: 0,
				input:          make([]float64, 2),
				output:         make([]float64, 2),
				isInit:         true,
			},
		},
		{
			name: "#2",
			got:  &NN{},
			want: &NN{
				Bias:   true,
				Hidden: []int{2},
				Weights: pkg.Float3Type{
					{
						{r, r, r},
						{r, r, r},
					},
					{
						{r, r, r},
					},
				},
				neuron: [][]*neuron{
					{
						&neuron{},
						&neuron{},
					},
					{
						&neuron{},
					},
				},
				lenInput:       2,
				lenOutput:      1,
				lastLayerIndex: 1,
				input:          make([]float64, 2),
				output:         make([]float64, 1),
				isInit:         true,
			},
		},
	}
	for _, tt := range tests {
		tt.got.Bias = tt.want.Bias
		tt.got.Hidden = tt.want.Hidden
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.initFromNew(tt.want.lenInput, tt.want.lenOutput); !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("initFromNew()\ngot:\t%v\nwant:\t%v", tt.got, tt.want)
			}
		})
	}
}

func TestNN_initFromWeight(t *testing.T) {
	tests := []struct {
		name string
		got  *NN
		want *NN
	}{
		{
			name: "#1",
			got:  &NN{},
			want: &NN{
				Hidden: []int{0},
				Weights: pkg.Float3Type{
					{
						{.1, .1},
						{.1, .1},
					},
				},
				neuron: [][]*neuron{
					{
						&neuron{},
						&neuron{},
					},
				},
				lenInput:       2,
				lenOutput:      2,
				lastLayerIndex: 0,
				input:          make([]float64, 2),
				output:         make([]float64, 2),
				isInit:         true,
			},
		},
		{
			name: "#2",
			got:  &NN{},
			want: &NN{
				Bias:   true,
				Hidden: []int{2},
				Weights: pkg.Float3Type{
					{
						{.1, .1, .1},
						{.1, .1, .1},
					},
					{
						{.1, .1, .1},
					},
				},
				neuron: [][]*neuron{
					{
						&neuron{},
						&neuron{},
					},
					{
						&neuron{},
					},
				},
				lenInput:       2,
				lenOutput:      1,
				lastLayerIndex: 1,
				input:          make([]float64, 2),
				output:         make([]float64, 1),
				isInit:         true,
			},
		},
	}
	for _, tt := range tests {
		tt.got.Weights = tt.want.Weights
		t.Run(tt.name, func(t *testing.T) {
			if tt.got.initFromWeight(); !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("initFromWeight()\ngot:\t%v\nwant:\t%v", tt.got, tt.want)
			}
		})
	}
}
