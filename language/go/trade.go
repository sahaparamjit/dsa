package main

import "fmt"

// Trade of stocks
type Trade struct {
	Symbol string // Ticker symbol of the traded stock
	Volume int // Number of stocks purchased from market
	Price float64 // Price of the position at the time of purchase
	Buy bool // true for buy , false for sell
}

// Value :- The value traded for the trade
func (t *Trade) Value() float64 {
	val := float64(t.Volume) * t.Price
	if t.Buy {
		val = -val
	}
	return val
}

func main() {
	msft := Trade{"MSFT", 3, 98.99, true}

	msft2 := Trade{
		Symbol : "MSFT", 
		Price : 98.99, 
		Volume : 3,
		Buy : true,
	}

	fmt.Printf("%+v\n", msft)
	fmt.Printf("%+v\n", msft2)

	fmt.Printf("%+v\n", msft.Value())
	fmt.Printf("%+v\n", msft2.Value())
}