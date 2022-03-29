package internal

import "fmt"

type HiGreeter struct {
	Name string
}

func (g HiGreeter) Greet() string {
	return fmt.Sprintf("Hi %s", g.Name)
}

func NewHiGreeter(name string) HiGreeter {
	return HiGreeter{
		Name: name,
	}
}
