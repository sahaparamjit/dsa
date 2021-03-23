package main

import "fmt"

func main() {
	mp := map[string]float64{
		"AMZN": 19.98,
		"MSFT": 59.98,
		"AAPL": 94.98,
	}

	for key, val := range mp {
		fmt.Printf("%v has stock price of %v\n", key, val)
	}

	for key := range mp {
		if mp[key] != 0 {
			fmt.Printf("%v has stock prices listed -> %v\n", key, mp[key])	
		}
	}

	exit, ok := mp["PYPL"]

	if !ok {
		fmt.Println("PYPL is not listed")
	}

	if exit==0 {
		fmt.Println("PYPL is not has zero exit code")
	}
}