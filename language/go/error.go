package main

import (
	"fmt"
	"math"
)

func eval(n float64) {
	sqrt, err := sqrt(n)

	if err == nil {
		fmt.Printf("%v is not a negative number and the square root of the number is %v\n", 2.0, sqrt)
	} else {
		fmt.Printf("%v \n", err)
	}
}

func main() {
	eval(2.0)
	eval(-2.0)


}


func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0.0, fmt.Errorf("Illegal Argument passed : Square root of a Negative number -> %v", a)
	}
	return math.Sqrt(a), nil
}