package shop

import (
	"testing"
)

func TestGet(t *testing.T) {
	shop, err := Get()
	if err != nil {
		t.Errorf("Get shop call failed: %v", err)
	}

	const expectedNumTreats = 4
	numTreats := len(shop.Treats)
	if numTreats != expectedNumTreats {
		t.Errorf("Expected %d treats, but got %d", expectedNumTreats, numTreats)
	}

	// Check first item
	testIndex := 0
	testTreat := shop.Treats[testIndex]
	expectedID := 1
	expectedName := "Brownie"
	expectedImageURL := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTHdr1eTXEMs68Dx-b_mZT0RpifEQ8so6A1unRsJlyJIPe0LUE2HQ"
	expectedPrice := float32(2.0)
	expectedBulkPriceAmount := 4
	expectedBulkPriceTotalPrice := float32(7.0)

	if testTreat.ID != expectedID {
		t.Errorf("Expected treats[%d] to have id %d, but got %d", testIndex, expectedID, testTreat.ID)
	}

	if testTreat.Name != expectedName {
		t.Errorf("Expected treats[%d] to have name '%s', but got '%s'", testIndex, expectedName, testTreat.Name)
	}

	if testTreat.ImageURL != expectedImageURL {
		t.Errorf("Expected treats[%d] to have imageURL '%s', but got '%s'", testIndex, expectedImageURL, testTreat.ImageURL)
	}

	if testTreat.Price != expectedPrice {
		t.Errorf("Expected treats[%d] to have price %f, but got %f", testIndex, expectedPrice, testTreat.Price)
	}

	if testTreat.BulkPricing.Amount != expectedBulkPriceAmount ||
		testTreat.BulkPricing.TotalPrice != expectedBulkPriceTotalPrice {
		t.Errorf("Expected treats[%d] bulkPricing to be %d for %f, but got %d for %f",
			testIndex, expectedBulkPriceAmount, expectedBulkPriceTotalPrice,
			testTreat.BulkPricing.Amount, testTreat.BulkPricing.TotalPrice)
	}

	// Check item without bulk pricing
	testIndexNoBulk := 3
	testTreatNoBulk := shop.Treats[testIndexNoBulk]
	if testTreatNoBulk.BulkPricing != nil {
		t.Errorf("Expected treats[%d] to have nil bulkPricing, but got %v", testIndexNoBulk, testTreatNoBulk.BulkPricing)
	}
}
