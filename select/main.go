package main

import (
	"fmt"
	"math/rand"
	"time"
)

func findSC(c chan int, name string) {
	scMapping := map[string]int{
		"James": 5,
		"Kevin": 10,
		"Rahul": 9,
	}

	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

	c <- scMapping[name]
}

func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	name := "Kevin"

	go findSC(c1, name)
	go findSC(c2, name)

	select {
	case sc := <-c1:
		fmt.Println(name, "has a security clearance of", sc, "found in server 1")
	case sc := <-c2:
		fmt.Println(name, "has a security clearance of", sc, "found in server 2")
	case <-time.After(3 * time.Second):
		fmt.Println("Search time out")
	}

}
