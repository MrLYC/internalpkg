package main

import (
	"fmt"
	"os"

	"github.com/mrlyc/internalpkg/greeting"
	// "github.com/mrlyc/internalpkg/greeting/internal"
)

func greet(action, name string) string {
	var greeter greeting.Greeter
	switch os.Args[1] {
	case "hi":
		greeter = greeting.NewHiGreeter(name)
	case "hello":
		greeter = greeting.NewHelloGreeter(name)
	default:
		return ""
	}

	return greeter.Greet()
}

// func internalGreet(action, name string) string {
// 	var greeter greeting.Greeter
// 	switch os.Args[1] {
// 	case "hi":
// 		greeter = internal.NewHiGreeter(name)
// 	case "hello":
// 		greeter = internal.NewHelloGreeter(name)
// 	default:
// 		return ""
// 	}

// 	return greeter.Greet()
// }

func main() {
	if len(os.Args) == 0 {
		fmt.Println("name is required")
		return
	}

	fmt.Println(greet(os.Args[1], "world"))
}
