package api

import "encoding/json"
import "goanda/models"

func parseAccountOpenPositions(msg *[]byte) (models.AccountPositions, error) {
	var p models.AccountPositions
	err := json.Unmarshal(*msg, &p)
	return p, err
}
