package main

import (
	"fmt"
	"time"
)

func Ping(c chan string) {
	for {
		c <- "PING"
	}
}

func Pong(c chan string) {
	for {
		c <- "PONG"
	}
}

func PrintChan(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	pp := make(chan string)

	// go func() {
	// for {
	// time.Sleep(1 * time.Second)
	// pp <- "PING"
	// time.Sleep(1 * time.Second)
	// pp <- "PONG"
	// }}()

	fmt.Println("PRESS ENTER/RETURN TO STOP")
	// time.Sleep(3 * time.Second)
	// for {
	// 	p1 := <-pp
	// 	fmt.Println(p1)
	// }
	go Ping(pp)
	go PrintChan(pp)
	go Pong(pp)

	var keypress string
	fmt.Scanln(&keypress)
}
