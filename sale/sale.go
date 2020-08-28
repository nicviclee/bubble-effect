package sale

import (
	"time"

	"github.com/nicviclee/bubble-effect/shop"
)

// GetSalesForDate returns a map of shop item to a function that applies a sale
// when applicable for the given date. The sale function takes in the item's
// regular price and the amount being purchased; it returns the total price
// after applying the sale and the remaining item count for items not included
// in the sale discount
func GetSalesForDate(date time.Time) map[shop.Item]func(regularPrice float32, amount int) (salePriceTotal float32, remainingUnits int) {
	sales := make(map[shop.Item]func(float32, int) (float32, int))

	// Every Friday
	if date.Weekday() == time.Friday {
		sales[shop.Cookie] = applyFridayCookieSale
	}

	// Every October 1st
	if date.Month() == time.October && date.Day() == 1 {
		sales[shop.KeyLimeCheesecake] = applyOctoberFirstCheesecakeSale
	}

	// Every Tuesday
	if date.Weekday() == time.Tuesday {
		sales[shop.MiniGingerbreadDonut] = applyTuesdayDonutSale
	}

	return sales
}

func applyFridayCookieSale(regularPrice float32, amount int) (salePriceTotal float32, remainingUnits int) {
	// 8 cookies for $6
	saleUnits := amount / 8
	salePrice := float32(saleUnits) * 6.00

	remainingUnits = amount % 8

	return salePrice, remainingUnits
}

func applyOctoberFirstCheesecakeSale(regularPrice float32, amount int) (salePriceTotal float32, remainingUnits int) {
	// 25% off keylime cheesecake
	salePrice := float32(amount) * regularPrice * 0.75
	return salePrice, 0
}

func applyTuesdayDonutSale(regularPrice float32, amount int) (salePriceTotal float32, remainingUnits int) {
	// Two for one mini gingerbread donuts
	saleUnits := amount / 2
	salePrice := float32(saleUnits) * regularPrice

	remainingUnits = amount % 2

	return salePrice, remainingUnits
}
