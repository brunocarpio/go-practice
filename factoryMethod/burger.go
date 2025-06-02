package main

import "fmt"

type IBurger interface {
	Prepare()
}

type Burger struct {
	name string
}

func (b *Burger) Prepare() {
	fmt.Println("Preparing burger,", b.name)
}

type BeefBurger struct {
	Burger
}

func newBeefBurger() IBurger {
	return &BeefBurger{
		Burger: Burger{
			name: "Beef",
		},
	}
}

type VeggieBurger struct {
	Burger
}

func newVeggieBurger() IBurger {
	return &VeggieBurger{
		Burger: Burger{
			name: "Veggie",
		},
	}
}
