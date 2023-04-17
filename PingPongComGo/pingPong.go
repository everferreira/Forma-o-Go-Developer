package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
		time.Sleep(time.Second)
	}
}

func pong(c chan string) {
	for {
		c <- "pong"
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)

	for {
		fmt.Println(<-c)
	}
}
