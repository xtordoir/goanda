package models

// AccountPositions are the Positions associated with an account
type AccountPositions struct {
	LastTransactionID string     `json:"lastTransactionID"`
	Positions         []Position `json:"positions"`
}
