package models

// Position Definitions

// Order is an order definition
type Order struct {
	Units        int64  `json:"units,string"`
	Instrument   string `json:"instrument"`
	TimeInForce  string `json:"timeInForce"`
	Type         string `json:"type"`
	PositionFill string `json:"positionFill"`
}

// OrderRequest is an order payload
type OrderRequest struct {
	Order Order `json:"order"`
}

// MakeMarketOrder creates a martket Order
func MakeMarketOrder(instrument string, units int64) Order {
	return Order{
		Units:        units,
		Instrument:   instrument,
		TimeInForce:  "FOK",
		Type:         "MARKET",
		PositionFill: "DEFAULT",
	}
}
