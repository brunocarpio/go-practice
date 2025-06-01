package main

import (
	"fmt"
	"math/rand"
	"time"
)

type customRand struct { // outer type
	*rand.Rand // inner type
	count      int
}

func newCustomRand(i int64) *customRand {
	return &customRand{
		Rand:  rand.New(rand.NewSource(i)),
		count: 0,
	}
}

func (cr *customRand) RandRange(min, max int) int {
	cr.count++
	return cr.Rand.Intn(max-min) + min
}

func (cr *customRand) Intn(n int) int {
	fmt.Printf("\"Outer is Intn called\": %v\n", "Outer is Intn called")
	cr.count++
	return cr.Rand.Intn(n)
}

func main() {
	cr := newCustomRand(time.Now().UnixNano())
	fmt.Printf("cr.RandRange(5, 30): %v\n", cr.RandRange(5, 30))
	fmt.Printf("cr.Intn(10): %v\n", cr.Intn(10))
	fmt.Printf("cr.count: %v\n", cr.count)

	// var r *rand.Rand = cr
	// cr = &rand.Rand{}
}
