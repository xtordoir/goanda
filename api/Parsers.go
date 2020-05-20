package api

import "encoding/json"
import "github.com/xtordoir/goanda/models"

func parseAccountOpenPositions(msg *[]byte) (models.AccountPositions, error) {
	var p models.AccountPositions
	err := json.Unmarshal(*msg, &p)
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
