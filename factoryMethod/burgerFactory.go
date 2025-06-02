package main

import "errors"

const (
	BEEF = iota
	VEGGIE
)

func createBurger(request int) (IBurger, error) {
	switch request {
	case BEEF:
		return newBeefBurger(), nil
	case VEGGIE:
		return newVeggieBurger(), nil
	default:
		return nil, errors.New("Invalid Burger type")
	}
}
