package models

// Position Definitions

// Position is a position in an account
type Position struct {
	Instrument string       `json:"instrument"`
	Long       PositionSide `json:"long"`
	Short      PositionSide `json:"short"`
}

// PositionSide is a Position for a single direction
type PositionSide struct {
	AveragePrice float64  `json:"averagePrice,string"`
	PL           float64  `json:"pl,string"`
	ResettablePL float64  `json:"resettablePL,string"`
	TradeIDs     []string `json:"tradeIDs"`
	Units        int64    `json:"units,string"`
	UnrealizedPL float64  `json:"unrealizedPL,string"`
}

// CalculatedPositionSide Not Implemented
