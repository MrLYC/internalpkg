package internal

import "fmt"

type HiGreeter struct {
	BaseGreeter
}

func (g HiGreeter) Greet() string {
	return fmt.Sprintf("Hi %s", g.Name())
}

func NewHiGreeter(name string) HiGreeter {
	return HiGreeter{
		BaseGreeter: NewBaseGreeter(name),
	}
}
