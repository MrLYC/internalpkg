package internal

import "fmt"

type HiGreeter struct {
	name string
}

func (g HiGreeter) Greet() string {
	return fmt.Sprintf("Hi %s", g.name)
}

func NewHiGreeter(name string) HiGreeter {
	return HiGreeter{
		name: name,
	}
}
