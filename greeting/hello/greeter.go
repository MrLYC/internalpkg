package hello

import "github.com/mrlyc/internalpkg/greeting/internal"

func New(name string) internal.MagicGreeter {
	return internal.New("Hello", name)
}
