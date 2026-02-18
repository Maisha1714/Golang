package main

import "fmt"

func main() {
	isthatgirl := false

	if isthatgirl {
		fmt.Println("Yes, she is!")
	} else {
		fmt.Println("No, she isn't.")
	}

	fmt.Println("Hello, World!")
	a := 2
	switch a {
	case 1:
		fmt.Println("One")

	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other number")
	}
}
