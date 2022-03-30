package internal

import "fmt"

type MagicGreeter struct {
	Action string
	Name   string
}

func (g MagicGreeter) Greet() string {
	return fmt.Sprintf("%s %s", g.Action, g.Name)
}

func New(action, name string) MagicGreeter {
	return MagicGreeter{
		Action: action,
		Name:   name,
	}
}
