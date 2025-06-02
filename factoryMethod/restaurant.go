package main

type IRestaurant interface {
	CreateBurger() IBurger
	OrderBurger() IBurger
}

type AbstractRestaurant struct{ IRestaurant }

func (res AbstractRestaurant) OrderBurger() IBurger {
	burger := res.CreateBurger()
	burger.Prepare()
	return burger
}

type BeefBurgerRestaurant struct {
	AbstractRestaurant
}

func (res BeefBurgerRestaurant) CreateBurger() IBurger {
	return newBeefBurger()
}

func NewBeefBurgerRestaurant() BeefBurgerRestaurant {
	res := BeefBurgerRestaurant{AbstractRestaurant{}}
	res.AbstractRestaurant.IRestaurant = res
	return res
}

type VeggieBurgerRestaurant struct {
	AbstractRestaurant
}

func (res VeggieBurgerRestaurant) CreateBurger() IBurger {
	return newVeggieBurger()
}

func NewVeggieBurgerRestaurant() VeggieBurgerRestaurant {
	res := VeggieBurgerRestaurant{AbstractRestaurant{}}
	res.AbstractRestaurant.IRestaurant = res
	return res
}
