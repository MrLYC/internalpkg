package greeting

import "github.com/mrlyc/internalpkg/greeting/internal"

type Greeter interface {
	Greet() string
}

func NewHiGreeter(name string) Greeter {
	return internal.NewHiGreeter(name)
}

func NewHelloGreeter(name string) Greeter {
	return internal.NewHelloGreeter(name)
}
