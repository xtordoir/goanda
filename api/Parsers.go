package api

import (
	"encoding/json"
	"errors"

	"github.com/xtordoir/goanda/models"
)

func parseAccountOpenPositions(msg *[]byte) (models.AccountPositions, error) {
	var p models.AccountPositions
	err := json.Unmarshal(*msg, &p)
	if err == nil && p.LastTransactionID == "" {
		return p, errors.New("No data: LastTransactionID empty")
	}
	return p, err
}

func parseAccountOrders(msg *[]byte) (models.AccountOrders, error) {
	var p models.AccountOrders
	err := json.Unmarshal(*msg, &p)
	if err == nil && p.LastTransactionID == "" {
		return p, errors.New("No data: LastTransactionID empty")
	}
	return p, err
}

func parseAccountPosition(msg *[]byte) (models.AccountPosition, error) {
	var p models.AccountPosition
	err := json.Unmarshal(*msg, &p)
	if err == nil && p.LastTransactionID == "" {
		return p, errors.New("No data: LastTransactionID empty")
	}
	return p, err
}

func parseClientPrice(msg *[]byte) (models.ClientPrice, error) {
	var p models.ClientPrice
	err := json.Unmarshal(*msg, &p)
	return p, err
}

func parsePrices(msg *[]byte) (models.Prices, error) {
	var p models.Prices
	err := json.Unmarshal(*msg, &p)
	return p, err
}

func parseCandles(msg *[]byte) (models.Candles, error) {
	var p models.Candles
	err := json.Unmarshal(*msg, &p)
	return p, err
}

func parsePositionBook(msg *[]byte) (models.PositionBook, error) {
	var p models.PositionBookResponse
	err := json.Unmarshal(*msg, &p)
	return p.PositionBook, err
}

func parseAccounts(msg *[]byte) (models.Accounts, error) {
	var acc models.Accounts
	err := json.Unmarshal(*msg, &acc)

	return acc, err
}
