package main

func main() {
	beefRestaurant := NewBeefBurgerRestaurant()
	beefRestaurant.OrderBurger()

	veggieRestaurant := NewVeggieBurgerRestaurant()
	veggieRestaurant.OrderBurger()
}
