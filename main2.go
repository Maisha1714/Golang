package main

import "fmt"

func add(a int, b int) int {

	sum := a + b
	fmt.Println(sum)
	return sum
}

func main() {
	x := 20
	y := 30
	add(x, y)
	add(100, 99)
}
