package greeting

import (
	"github.com/mrlyc/internalpkg/greeting/hello"
	"github.com/mrlyc/internalpkg/greeting/hi"
	"github.com/mrlyc/internalpkg/greeting/internal"
)

type Greeter interface {
	Greet() string
}

func NewHiGreeter(name string) Greeter {
	return hi.New(name)
}

func NewHelloGreeter(name string) Greeter {
	return hello.New(name)
}

func NewMagicGreeter(action, name string) Greeter {
	return internal.New(action, name)
}
