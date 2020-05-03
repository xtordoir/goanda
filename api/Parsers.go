package api

import "encoding/json"
import "github.com/xtordoir/goanda/models"

func parseAccountOpenPositions(msg *[]byte) (models.AccountPositions, error) {
	var p models.AccountPositions
	err := json.Unmarshal(*msg, &p)
	return p, err
}
