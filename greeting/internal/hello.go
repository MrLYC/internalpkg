package internal

import "fmt"

type HelloGreeter struct {
	name string
}

func (g HelloGreeter) Greet() string {
	return fmt.Sprintf("Hello %s", g.name)
}

func NewHelloGreeter(name string) HelloGreeter {
	return HelloGreeter{
		name: name,
	}
}
