package cart

import (
	"fmt"
	"time"

	"github.com/nicviclee/bubble-effect/sale"
	"github.com/nicviclee/bubble-effect/shop"
)

// CalculateCartPrice returns the cart price for the given items and amounts,
// taking into account active sales, bulk pricing, and regular pricing
func CalculateCartPrice(date time.Time, itemsWithAmounts map[shop.Item]int) (float32, error) {
	// Refreshes shop data in case items or prices change
	// TODO: Consider caching to reduce the number of calls
	shop, err := shop.Get()
	if err != nil {
		return 0, fmt.Errorf("could not get shop data : %v", err)
	}

	total := float32(0)
	sales := sale.GetSalesForDate(date)

	for item, amount := range itemsWithAmounts {
		treatDetails := shop.Treats[item]

		remainingUnits := amount

		// Apply sale pricing
		var salePriceTotal float32
		applySalePricing, ok := sales[item]
		if ok {
			// There is an applicable sale
			salePriceTotal, remainingUnits = applySalePricing(treatDetails.Price, amount)
			total += salePriceTotal
		}

		// Apply bulk pricing to remaining count
		bulkPriceTotal, remainingUnits := applyBulkPricing(treatDetails, remainingUnits)
		total += bulkPriceTotal

		// Apply regular pricing to remaining count
		total += treatDetails.Price * float32(remainingUnits)
	}

	return total, nil
}

func applyBulkPricing(treatDetails shop.Treat, amount int) (totalBulkPrice float32, remainingUnits int) {
	if treatDetails.BulkPricing == nil {
		return 0, amount
	}

	bulkPriceUnits := amount / treatDetails.BulkPricing.Amount
	remainingUnits = amount % treatDetails.BulkPricing.Amount

	totalBulkPrice = treatDetails.BulkPricing.TotalPrice * float32(bulkPriceUnits)

	return totalBulkPrice, remainingUnits
}
