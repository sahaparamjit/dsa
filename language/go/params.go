package main

import "fmt"


func doubleAt(values []int, i int) {
	values[i]*=2
}


func main() {
	arr := []int{1 ,2 , 3, 4}
	doubleAt(arr, 2)
	for _, val := range arr {
		fmt.Printf("%v\n", val)
	}
	val := 2 
	fmt.Printf("Before %v\n", val)
	doublePtr(&val)
	fmt.Printf("After %v\n", val)
}

func doublePtr(val *int) {
	*val *= 2
}