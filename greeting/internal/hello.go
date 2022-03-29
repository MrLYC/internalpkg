package internal

import "fmt"

type HelloGreeter struct {
	Name string
}

func (g HelloGreeter) Greet() string {
	return fmt.Sprintf("Hello %s", g.Name)
}

func NewHelloGreeter(name string) HelloGreeter {
	return HelloGreeter{
		Name: name,
	}
}
