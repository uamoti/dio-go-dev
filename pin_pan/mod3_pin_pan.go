package main

import "fmt"

func main() {
    fmt.Println("Multiples of 3 between 1 and 100")
    for i := 1; i < 101; i++ {
    	if i % 3 == 0 {
	   fmt.Printf("%v ", i)
	}
    }
    fmt.Println()

    fmt.Println("Pin and Pan")
    for i := 1; i < 101; i++ {
    	if i % 3 == 0 {
	   fmt.Print("Pin ")
	} else if i % 5 == 0 {
	  fmt.Print("Pan ")
	} else {
	  fmt.Printf("%v ", i)
	}
    }
    fmt.Println()
}
