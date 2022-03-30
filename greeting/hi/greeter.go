package hi

import "github.com/mrlyc/internalpkg/greeting/internal"

func New(name string) internal.MagicGreeter {
	return internal.New("Hi", name)
}
