package main

import "fmt"

func main() {
    var tempCels float32

    fmt.Println("*** CELSIUS TO KELVIN CONVERTER ***")
    fmt.Print("Temperature in Celcius: ")
    fmt.Scan(&tempCels)
    fmt.Println(tempCels, "degrees Celsius equals", tempCels + 273, "degrees Kelvin.")
}
