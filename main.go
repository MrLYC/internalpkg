package main

import (
	"fmt"
	"os"

	"github.com/mrlyc/internalpkg/greeting"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <action> <name>", os.Args[0])
		os.Exit(1)
	}

	var greeter greeting.Greeter
	action := os.Args[1]
	name := os.Args[2]
	switch action {
	case "hi":
		greeter = greeting.NewHiGreeter(name)
	case "hello":
		greeter = greeting.NewHelloGreeter(name)
	default:
		greeter = greeting.NewMagicGreeter(action, name)
	}

	fmt.Println(greeter.Greet())
}
