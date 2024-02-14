package models

import "fmt"
// Position Definitions

// AccountOrder is an Order in the Account
type AccountOrder struct {
	Id					 int64  `json:"id,string"`
	Units        Unit   `json:"units,string"`
	Price        float64 `json:"price,string"`
	Instrument   string `json:"instrument"`
}

// Order is an order definition
type Order struct {
	Units        Unit  `json:"units,string"`
	Instrument   string `json:"instrument"`
	TimeInForce  string `json:"timeInForce"`
	Type         string `json:"type"`
	PositionFill string `json:"positionFill"`
}

type LimitOrder struct {
	Price        string  `json:"price,string"`
	Units        string  `json:"units,string"`
	Instrument   string `json:"instrument"`
	TimeInForce  string `json:"timeInForce"`
	Type         string `json:"type"`
	PositionFill string `json:"positionFill"`
}

func NewLimitOrder(instrument string, units Unit, unitsDecimals int, price float64, priceDecimals int) LimitOrder {
	return LimitOrder {
		Price: fmt.Sprintf("%.[1]*f", priceDecimals, price),
		Units: fmt.Sprintf("%.[1]*f", unitsDecimals, float64(units)),
		Instrument:   instrument,
		TimeInForce:  "GTC",
		Type:         "LIMIT",
		PositionFill: "DEFAULT",
	}
}
// OrderRequest is an order payload
type LimitOrderRequest struct {
	Order LimitOrder `json:"order"`
}

// OrderRequest is an order payload
type OrderRequest struct {
	Order Order `json:"order"`
}

// MakeMarketOrder creates a martket Order
func MakeMarketOrder(instrument string, units Unit) Order {
	return Order{
		Units:        units,
		Instrument:   instrument,
		TimeInForce:  "FOK",
		Type:         "MARKET",
		PositionFill: "DEFAULT",
	}
}
