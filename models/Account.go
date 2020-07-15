package models

// AccountPositions are the Positions associated with an account
type AccountPositions struct {
	LastTransactionID string     `json:"lastTransactionID"`
	Positions         []Position `json:"positions"`
}

// AccountPosition is a single Position associated with an account
type AccountPosition struct {
	LastTransactionID string   `json:"lastTransactionID"`
	Position          Position `json:"position"`
}
