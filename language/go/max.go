package main

import "fmt"

func main() {

	intArr := []int{16, 187, 90, 2, 4}

	max := intArr[0]

	for _, val := range intArr[1:] {
		if max < val {
			max = val
		}
	}

	fmt.Printf("Max value in %v is : %v\n", intArr, max)
}