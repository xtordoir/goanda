package models

// Account implementation
type Account struct {
	ID string `json:"id"`
	// only ID is implemented at this stage
}

// Accounts is the structure returned by GET Accounts endpoint
type Accounts struct {
	Accounts []Account `json:"accounts"`
}

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
