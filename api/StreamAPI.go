package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/xtordoir/goanda/models"
)

// StreamAPI is an api instance with a context to call endpoints
type StreamAPI struct {
	context Context
}

type priceProcessor func(p *models.ClientPrice)
type heartbeatProcessor func(p *models.PricingHeartbeat)

// PricingStream starts a stream of prices
func (streamApi *StreamAPI) PricingStream(instruments []string, pprocessor priceProcessor, hbProcessor heartbeatProcessor) {

	var p models.ClientPrice
	var t models.Tick

	url := streamApi.context.ApiURL + "/" + streamApi.context.Account + "/pricing/stream"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url+"?instruments="+strings.Join(instruments, ","), nil)
	req.Header.Add("Authorization", "Bearer "+streamApi.context.Token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		reader := bufio.NewReader(response.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				panic(err)
			}
			json.Unmarshal([]byte(line), &p)
			if p.Type == "HEARTBEAT" {
				h := models.PricingHeartbeat{
					Type: p.Type,
					Time: p.Time,
				}
				go hbProcessor(&h)
				//fmt.Println(h)
			} else {
				go pprocessor(&p)
				t = models.ClientPrice2Tick(&p)
				fmt.Println(t)
			}
		}
	}
	//return nil
}
