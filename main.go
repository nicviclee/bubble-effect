package main

import (
	"fmt"
	"time"

	"github.com/nicviclee/bubble-effect/cart"
	"github.com/nicviclee/bubble-effect/shop"
)

var (
	noSaleDate           = time.Date(2020, time.August, 27, 0, 0, 0, 0, time.UTC)
	octoberFirstSaleDate = time.Date(2021, time.October, 1, 0, 0, 0, 0, time.UTC)
)

func main() {
	fmt.Println("Welcome to the CAI Bakery")

	// Order #1
	price, err := cart.CalculateCartPrice(noSaleDate,
		map[shop.Item]int{
			shop.Cookie:            1,
			shop.Brownie:           4,
			shop.KeyLimeCheesecake: 1,
		})

	printPriceOrError("Cart #1", price, err) // $16.25

	// Order #2
	price, err = cart.CalculateCartPrice(noSaleDate,
		map[shop.Item]int{shop.Cookie: 8})

	printPriceOrError("Cart #2", price, err) // $8.50

	// Order #3
	price, err = cart.CalculateCartPrice(noSaleDate,
		map[shop.Item]int{
			shop.Cookie:               1,
			shop.Brownie:              1,
			shop.KeyLimeCheesecake:    1,
			shop.MiniGingerbreadDonut: 2,
		})

	printPriceOrError("Cart #3", price, err) // $12.25

	// Order #4 - cheesecake sale
	price, err = cart.CalculateCartPrice(octoberFirstSaleDate,
		map[shop.Item]int{
			shop.Cookie:            8,
			shop.KeyLimeCheesecake: 4,
		})

	printPriceOrError("Cart #4", price, err) // $30.00

	fmt.Println("Thanks for shopping!")
}

func printPriceOrError(name string, price float32, err error) {
	if err != nil {
		fmt.Printf("%s error: %v\n", name, err)
	} else {
		fmt.Printf("%s price: %0.2f\n", name, price)
	}
}
