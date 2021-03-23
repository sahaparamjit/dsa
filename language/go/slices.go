package main

import "fmt"

func main() {

	v := []string{"hello", "jillu", "moly"}

	fmt.Printf("%v slice of (type %T)\n", v, v)

	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}

	fmt.Println(v[:3])

	fmt.Println(v[1:2])

	for i, name := range v {
		fmt.Printf("%v -> %v\n", i, name)
	}

	v = append(v, "nice")
	for _, name := range v {
		fmt.Printf("-> %v\n", name)
	}

	
}