package models

type AccountDetails struct {
	LastTransactionID string  `json:"lastTransactionID"`
	Account           Account `json:"account"`
}

// Account implementation
type Account struct {
	ID string  `json:"id"`
	// only ID is implemented at this stage
	//LastTransactionID string     `json:"lastTransactionID"`
	// we cannot unmarshall these fields directly from GetAccount , not sure why
	Positions     []Position `json:"positions"`
	// TODO support all orders types...
	Orders        []Order `json:"orders"`

	Nav float64 `json:"NAV,string"`

	PositionValue float64 `json:"positionValue,string"`
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

type AccountOrders struct {
	LastTransactionID string               `json:"lastTransactionID"`
	Orders            []AccountOrder       `json:"orders"`
}
