package models

// Transactions Definitions

import (
	"time"
)

// Transaction is a Generic Transaction, just with an ID at the moment
type Transaction struct {
  ID string   `json:"id"`
  Type string `json:"type"`
}

// TransactionHeartbeat is a heartbeat to keep connection alive, containing LastTransactionID
type TransactionHeartbeat struct {
	Type string    `json:"type"`
  LastTransactionID string `json:"lastTransactionID"`
	Time time.Time `json:"time"`
}
