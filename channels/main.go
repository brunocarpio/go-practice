package main

import (
	"fmt"
)

func waitAndSay(c chan bool, s string) {
	if b := <-c; b {
		fmt.Println(s)
	}

	c <- true // use other channel as exit channel? a better approach
}

func main() {
	c := make(chan bool)

	go waitAndSay(c, "World")

	fmt.Println("Hello")

	c <- true

	hiFromBufferedChannel()

	<-c

	close(c)

	c1 := make(chan string)

	go sayHiNTimes(c1, 5)

	for s := range c1 {
		fmt.Println(s)
	}

	v, ok := <-c1
	fmt.Println("Is channel closed?", !ok, "value", v)

}

func hiFromBufferedChannel() {
	bc := make(chan string, 2)
	bc <- "hi"
	bc <- "neighbor"
	fmt.Println(<-bc, <-bc)
}

func sayHiNTimes(c chan string, n int) {
	for range n {
		c <- "hi"
	}
	close(c)
}
