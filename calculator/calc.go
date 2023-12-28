package main

import "fmt"

func Plus(arr ...float64) float64 {
	total := 0.0
	for _, val := range arr {
		total += val
	}
	return total
}

func Minus(arr ...float64) float64 {
	total := 0.0
	for i, val := range arr {
		if i == 0 {
			total = val
		} else {
			total -= val
		}
	}
	return total
}

func Mult(arr ...float64) float64 {
	total := 1.0
	for _, val := range arr {
		total *= val
	}
	return total
}

func Div(arr ...float64) float64 {
	var total float64
	for i, val := range arr {
		if i == 0 {
			total = val
		} else {
			total /= val
		}
	}
	return total
}

func main() {
	fmt.Println("Sum =", Plus(1, 2, 3))
	fmt.Println("Subtraction =", Minus(1, 2, 3))
	fmt.Println("Multiplication =", Mult(1, 2, 3))
	fmt.Println("Division =", Div(8, 2, 2))
}
