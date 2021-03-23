package main

import "fmt"


func main () {
	x := "I am a programmer"
	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}