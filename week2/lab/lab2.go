package lab

import (
	"fmt"
)

type Snack struct {
	Name  string
	Price int
}

type Drink struct {
	Name  string
	Price int
}

func (snack *Snack) getPrice() int {
	return snack.Price
}

func (drink *Drink) getPrice() int {
	return drink.Price
}

func (snack *Snack) Sale() {
	snack.Price = snack.Price * 9 / 10
}

func (drink *Drink) Sale() {
	drink.Price = drink.Price * 8 / 10
}

type product interface {
	Sale()
	getPrice() int
}

func SaleAndGetPrice(cart product) int {
	cart.Sale()
	return cart.getPrice()
}

func Lab2() {
	chips := Snack{"Pringles", 4000}
	cracker := Snack{"Ace", 2500}

	soda := Drink{"Sprite", 1800}
	coffee := Drink{"TOP", 2700}

	var total int = 0
	total += SaleAndGetPrice(&chips)
	total += SaleAndGetPrice(&cracker)
	total += SaleAndGetPrice(&soda)
	total += SaleAndGetPrice(&coffee)

	fmt.Println(total)
}
