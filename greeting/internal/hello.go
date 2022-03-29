package internal

import "fmt"

type HelloGreeter struct {
	BaseGreeter
}

func (g HelloGreeter) Greet() string {
	return fmt.Sprintf("Hello %s", g.Name())
}

func NewHelloGreeter(name string) HelloGreeter {
	return HelloGreeter{
		BaseGreeter: NewBaseGreeter(name),
	}
}
