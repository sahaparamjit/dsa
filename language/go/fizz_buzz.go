package main

import "fmt"

func main() {

	for i := 1; i<110; i++ {
		if ( i%5== 0 && i%3==0 ) {
			fmt.Println("FizzBuzz")
		} else if ( i%5 == 0 ) {
			fmt.Println("Buzz")
		} else if ( i%3==0 ){
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

}