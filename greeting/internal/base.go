package internal

type BaseGreeter struct {
	name string
}

func (g BaseGreeter) Name() string {
	return g.name
}

func NewBaseGreeter(name string) BaseGreeter {
	return BaseGreeter{
		name: name,
	}
}
