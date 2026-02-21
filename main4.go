package main

import "fmt"

// Multiple function
func add(a int, b int) int {

	sum := a + b
	return sum
}

func getnumbers(c int, d int) (int, int) {

	sub := c - d

	mul := c * d

	return sub, mul
}

func main() {
	x := 20
	y := 30

	az := 100
	ma := 10
	result := add(x, y)
	result2, result3 := getnumbers(az, ma)
	fmt.Println(result)
	fmt.Println(result2)
	fmt.Println(result3)
}
