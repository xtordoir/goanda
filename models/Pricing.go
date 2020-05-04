package models

// Pricing Definitions

import (
	"time"
)

// ClientPrice is the price of an instrument
type ClientPrice struct {
	Instrument string        `json:"instrument"`
	Type       string        `json:"type"`
	Time       time.Time     `json:"time"`
	Bids       []PriceBucket `json:"bids"`
	Asks       []PriceBucket `json:"asks"`
}

// PricingHeartbeat is a heartbeat to keep connection alive
type PricingHeartbeat struct {
	Type string    `json:"type"`
	Time time.Time `json:"time"`
}

// Prices is the object response from GetPricing call
type Prices struct {
	Prices []ClientPrice `json:"prices"`
}
