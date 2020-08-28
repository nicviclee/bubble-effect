package sale

import (
	"testing"
	"time"

	"github.com/nicviclee/bubble-effect/shop"
)

var (
	noSaleDate                 = time.Date(2020, time.August, 27, 0, 0, 0, 0, time.UTC)
	fridayCookieDate           = time.Date(2020, time.August, 28, 0, 0, 0, 0, time.UTC)
	octoberFirstCheesecakeDate = time.Date(2021, time.October, 1, 0, 0, 0, 0, time.UTC)
	tuesdayDonutDate           = time.Date(2020, time.November, 24, 0, 0, 0, 0, time.UTC)
)

func TestGetSalesForDate_NoSales(t *testing.T) {
	sales := GetSalesForDate(noSaleDate)
	if len(sales) > 0 {
		t.Error("Expected no sales")
	}
}

func TestGetSalesForDate_FridaySale(t *testing.T) {
	sales := GetSalesForDate(fridayCookieDate)
	applyFridaySale, ok := sales[shop.Cookie]
	if !ok {
		t.Error("Expected Friday cookie sale")
	}

	var tests = []struct {
		regularPrice           float32
		amount                 int
		expectedSalePriceTotal float32
		expectedRemainingUnits int
	}{
		{regularPrice: 1.25, amount: 0, expectedSalePriceTotal: 0.00, expectedRemainingUnits: 0},
		{regularPrice: 1.25, amount: 7, expectedSalePriceTotal: 0.00, expectedRemainingUnits: 7},
		{regularPrice: 1.25, amount: 8, expectedSalePriceTotal: 6.00, expectedRemainingUnits: 0},
		{regularPrice: 1.25, amount: 10, expectedSalePriceTotal: 6.00, expectedRemainingUnits: 2},
	}

	for _, test := range tests {
		salePriceTotal, remainingUnits := applyFridaySale(test.regularPrice, test.amount)
		if salePriceTotal != test.expectedSalePriceTotal ||
			remainingUnits != test.expectedRemainingUnits {
			t.Errorf("applyFridaySale(%0.2f, %d) = %0.2f, %d",
				test.regularPrice, test.amount, salePriceTotal, remainingUnits)
		}
	}
}

func TestGetSalesForDate_OctoberFirstSale(t *testing.T) {
	sales := GetSalesForDate(octoberFirstCheesecakeDate)
	applyOctoberFirstSale, ok := sales[shop.KeyLimeCheesecake]
	if !ok {
		t.Error("Expected October 1st keylime cheesecake sale")
	}

	var tests = []struct {
		regularPrice           float32
		amount                 int
		expectedSalePriceTotal float32
		expectedRemainingUnits int
	}{
		{regularPrice: 8.00, amount: 0, expectedSalePriceTotal: 0.00, expectedRemainingUnits: 0},
		{regularPrice: 8.00, amount: 2, expectedSalePriceTotal: 12.00, expectedRemainingUnits: 0},
		{regularPrice: 8.00, amount: 8, expectedSalePriceTotal: 48.00, expectedRemainingUnits: 0},
		{regularPrice: 8.00, amount: 20, expectedSalePriceTotal: 120.00, expectedRemainingUnits: 0},
	}

	for _, test := range tests {
		salePriceTotal, remainingUnits := applyOctoberFirstSale(test.regularPrice, test.amount)
		if salePriceTotal != test.expectedSalePriceTotal ||
			remainingUnits != test.expectedRemainingUnits {
			t.Errorf("applyOctoberFirstSale(%0.2f, %d) = %0.2f, %d",
				test.regularPrice, test.amount, salePriceTotal, remainingUnits)
		}
	}
}

func TestGetSalesForDate_TuesdaySale(t *testing.T) {
	sales := GetSalesForDate(tuesdayDonutDate)
	applyTuesdaySale, ok := sales[shop.MiniGingerbreadDonut]
	if !ok {
		t.Error("Expected Tuesday donut sale")
	}

	var tests = []struct {
		regularPrice           float32
		amount                 int
		expectedSalePriceTotal float32
		expectedRemainingUnits int
	}{
		{regularPrice: 0.50, amount: 0, expectedSalePriceTotal: 0.00, expectedRemainingUnits: 0},
		{regularPrice: 0.50, amount: 1, expectedSalePriceTotal: 0.00, expectedRemainingUnits: 1},
		{regularPrice: 0.50, amount: 3, expectedSalePriceTotal: 0.50, expectedRemainingUnits: 1},
		{regularPrice: 0.50, amount: 4, expectedSalePriceTotal: 1.00, expectedRemainingUnits: 0},
		{regularPrice: 0.50, amount: 13, expectedSalePriceTotal: 3.00, expectedRemainingUnits: 1},
	}

	for _, test := range tests {
		salePriceTotal, remainingUnits := applyTuesdaySale(test.regularPrice, test.amount)
		if salePriceTotal != test.expectedSalePriceTotal ||
			remainingUnits != test.expectedRemainingUnits {
			t.Errorf("applyTuesdaySale(%0.2f, %d) = %0.2f, %d",
				test.regularPrice, test.amount, salePriceTotal, remainingUnits)
		}
	}
}
