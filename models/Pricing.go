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

// Candles is the object returns by the GetCandles call
type Candles struct {
	Instrument  string        `json:"instrument"`
	Granularity string        `json:"granularity"`
	Candles     []CandleStick `json:"candles"`
}

// CandleStick is the structure for a single Candle data
type CandleStick struct {
	Time     time.Time       `json:"time"`
	Bid      CandleStickData `json:"bid"`
	Ask      CandleStickData `json:"ask"`
	Mid      CandleStickData `json:"mid"`
	Volume   int             `json:"volume"`
	Complete bool            `json:"complete"`
}

// CandleStickData is the actiual OHLC prices
type CandleStickData struct {
	O float64 `json:"o,string"`
	H float64 `json:"h,string"`
	L float64 `json:"l,string"`
	C float64 `json:"c,string"`
}
