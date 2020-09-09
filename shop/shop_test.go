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
	numTreats := len(shop)
	if numTreats != expectedNumTreats {
		t.Errorf("Expected %d treats, but got %d", expectedNumTreats, numTreats)
	}

	// Check first item
	testID := 1
	testTreat := shop[Item(testID)]
	expectedID := 1
	expectedName := "Brownie"
	expectedImageURL := "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTHdr1eTXEMs68Dx-b_mZT0RpifEQ8so6A1unRsJlyJIPe0LUE2HQ"
	expectedPrice := float32(2.0)
	expectedBulkPriceAmount := 4
	expectedBulkPriceTotalPrice := float32(7.0)

	if testTreat.ID != expectedID {
		t.Errorf("Expected treats[%d] to have id %d, but got %d", testID, expectedID, testTreat.ID)
	}

	if testTreat.Name != expectedName {
		t.Errorf("Expected treats[%d] to have name '%s', but got '%s'", testID, expectedName, testTreat.Name)
	}

	if testTreat.ImageURL != expectedImageURL {
		t.Errorf("Expected treats[%d] to have imageURL '%s', but got '%s'", testID, expectedImageURL, testTreat.ImageURL)
	}

	if testTreat.Price != expectedPrice {
		t.Errorf("Expected treats[%d] to have price %f, but got %f", testID, expectedPrice, testTreat.Price)
	}

	if testTreat.BulkPricing.Amount != expectedBulkPriceAmount ||
		testTreat.BulkPricing.TotalPrice != expectedBulkPriceTotalPrice {
		t.Errorf("Expected treats[%d] bulkPricing to be %d for %f, but got %d for %f",
			testID, expectedBulkPriceAmount, expectedBulkPriceTotalPrice,
			testTreat.BulkPricing.Amount, testTreat.BulkPricing.TotalPrice)
	}

	// Check item without bulk pricing
	testIDNoBulk := 4
	testTreatNoBulk := shop[Item(testIDNoBulk)]
	if testTreatNoBulk.BulkPricing != nil {
		t.Errorf("Expected treats[%d] to have nil bulkPricing, but got %v", testIDNoBulk, testTreatNoBulk.BulkPricing)
	}
}
