package shop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"runtime"
)

// Item is the type for an item in the store
type Item int

// Item types available in the store
const (
	Brownie Item = iota
	KeyLimeCheesecake
	Cookie
	MiniGingerbreadDonut
)

// String returns the string representation of the item type
func (i Item) String() string {
	return []string{"Brownie", "Key Lime Cheesecake", "Cookie", "Mini Gingerbread Donut"}[i]
}

// Shop is a collection of treats
type Shop struct {
	Treats []Treat `json:"treats"`
}

// Treat describes a treat
type Treat struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	ImageURL    string       `json:"imageURL"`
	Price       float32      `json:"price"` // TODO: Represent price as cents (int)
	BulkPricing *BulkPricing `json:"bulkPricing"`
}

// BulkPricing describes bulk pricing for a treat
type BulkPricing struct {
	Amount     int     `json:"amount"`
	TotalPrice float32 `json:"totalPrice"`
}

const productsFilename = "products-data.json"

// Get retrieves shop data or returns an error
func Get() (*Shop, error) {
	_, currentPath, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("error getting current path")
	}

	productsPath := path.Join(path.Dir(currentPath), productsFilename)
	content, err := ioutil.ReadFile(productsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %v", productsPath, err)
	}

	var shopInventory Shop
	err = json.Unmarshal(content, &shopInventory)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &shopInventory, nil
}
