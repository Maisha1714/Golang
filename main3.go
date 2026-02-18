package main

import "fmt"

func add(a int, b int) int {

	sum := a + b
	return sum
}

func main() {
	x := 20
	y := 30
	result := add(x, y)
	fmt.Println(result)
}
