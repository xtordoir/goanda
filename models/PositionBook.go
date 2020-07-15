package models

import "time"

// PositionBookResponse is the position book for an instrrument
type PositionBookResponse struct {
	PositionBook PositionBook `json:"positionBook"`
}

// PositionBook is the position book for an instrrument
type PositionBook struct {
	Instrument  string    `json:"instrument"`
	Time        time.Time `json:"time"`
	Price       float64   `json:"price,string"`
	BucketWidth float64   `json:"bucketWidth,string"`
	Buckets     []Bucket  `json:"buckets"`
}

// Bucket is for the position book at a price
type Bucket struct {
	Price             float64 `json:"price,string"`
	LongCountPercent  float64 `json:"longCountPercent,string"`
	ShortCountPercent float64 `json:"shortCountPercent,string"`
}
