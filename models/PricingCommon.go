package models

// Pricing Common Definitions

// PriceValue is not defined, we use float64 at the moment
// TODO implement proper decimal numbers

// PriceBucket is a type for Bids or Asks
type PriceBucket struct {
	Price     float64 `json:"price,string"`
	Liquidity int     `json:"liquidity"`
}
