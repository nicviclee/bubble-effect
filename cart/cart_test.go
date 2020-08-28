package cart

import (
	"testing"
	"time"

	"github.com/nicviclee/bubble-effect/shop"
)

var nonSaleDate = time.Date(2020, time.August, 27, 0, 0, 0, 0, time.UTC)

func TestShoppingCart(t *testing.T) {
	var tests = []struct {
		inputDate     time.Time
		inputItems    map[shop.Item]int
		expectedPrice float32
	}{
		// No sales
		{
			inputDate: nonSaleDate,
			inputItems: map[shop.Item]int{
				shop.Cookie:            1,
				shop.Brownie:           4,
				shop.KeyLimeCheesecake: 1,
			},
			expectedPrice: float32(16.25)},

		{
			inputDate: nonSaleDate,
			inputItems: map[shop.Item]int{
				shop.Cookie: 8,
			},
			expectedPrice: float32(8.50)},

		{
			inputDate: nonSaleDate,
			inputItems: map[shop.Item]int{
				shop.Cookie:               1,
				shop.Brownie:              1,
				shop.KeyLimeCheesecake:    1,
				shop.MiniGingerbreadDonut: 2,
			},
			expectedPrice: float32(12.25)},

		// Sales
		{inputDate: time.Date(2021, 10, 1, 0, 0, 0, 0, time.UTC),
			inputItems: map[shop.Item]int{
				shop.Cookie:            8,
				shop.KeyLimeCheesecake: 4,
			},
			expectedPrice: float32(30.00)},
	}

	for _, test := range tests {
		price, err := CalculateCartPrice(test.inputDate, test.inputItems)
		if err != nil {
			t.Errorf("date: %v | calculateCartPrice(%v) error: %v", test.inputDate, test.inputItems, err)
		}
		if price != test.expectedPrice {
			t.Errorf("date: %v | calculateCartPrice(%v) = %f; expected %f", test.inputDate, test.inputItems, price, test.expectedPrice)
		}
	}
}
