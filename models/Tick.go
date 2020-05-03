package models

import "time"

// Tick is a bid/ask for an instrument at a given time
type Tick struct {
	Instrument string
	Time       time.Time
	Bid        float64
	Ask        float64
}

// ClientPrice2Tick converts a ClientPrice to a Tick, by taking the first Bid and Ask
func ClientPrice2Tick(price *ClientPrice) Tick {
	return Tick{
		Instrument: price.Instrument,
		Time:       price.Time,
		Bid:        (*price).Bids[0].Price,
		Ask:        (*price).Asks[0].Price,
	}
}
