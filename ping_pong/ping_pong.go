package main

import (
"fmt"
"time"
)

func main() {
     pp := make(chan string)

     go func() {
     for {
     time.Sleep(1 * time.Second)
     pp <- "PING"
     time.Sleep(1 * time.Second)
     pp <- "PONG"
     }}()

     fmt.Println("PRESS CTR-C TO STOP")
     time.Sleep(3 * time.Second)
     for {
     p1 := <- pp
     fmt.Println(p1)
     }
}

