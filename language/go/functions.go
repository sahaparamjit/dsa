package main

import "fmt"

func add(a int, b int) int {
	return a+b
}

func mod(a int, b int) (int, int) {
	return a % b, a/b
}

func main() {
	mod, div := mod(4,3)

	fmt.Printf("%v is mod and %v is the division of %v & %v\n", mod, div, 4, 3)
}