package main

import "fmt"

func main() {
    var tempCelc float32

    fmt.Println("*** CELCIUS TO KELVIN CONVERTER ***")
    fmt.Print("Temperature in Celcius: ")
    fmt.Scan(&tempCelc)
    fmt.Println(tempCelc, "degrees Celcius equals", tempCelc + 273, "degrees Kelvin.")
}
