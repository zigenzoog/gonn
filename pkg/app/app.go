package app

import "github.com/zigenzoog/gonn/pkg"

// Declare conformity with Application interface
var _ Application = (*app)(nil)

type Application interface {
	pkg.Controller
}

type app struct {
}

// App returns a new application instance with the default parameters
func App() *app {
	return &app{}
}