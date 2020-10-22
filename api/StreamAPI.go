package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/xtordoir/goanda/models"
)

// StreamAPI is an api instance with a context to call endpoints
type StreamAPI struct {
	context Context
}

type priceProcessor func(p *models.ClientPrice)
type heartbeatProcessor func(p *models.PricingHeartbeat)

// TickStream starts a stream of ticks, hiding the Prices structs which are autoRestarted
func (streamApi *StreamAPI) TickStream(instruments []string, tchan chan models.Tick, hchan chan models.PricingHeartbeat) {
	pchan := make(chan models.ClientPrice)
	// AutoRestart for PricingStream
	go autoRestart("PricingStream", 0, func() { streamApi.PricingStream(instruments, pchan, hchan) })

	fmt.Println("Starting loop on Prices")
	for {
		price := <-pchan
		tchan <- models.ClientPrice2Tick(&price)
	}
}

// AutoRestart for the PricingStream function as connection reset can result in panic
func autoRestart(name string, nPanics int64, f func()) {
	defer func() {
		if v := recover(); v != nil {
			// A panic is detected.
			log.Printf("%s is crashed. Panic #%d. Restarting in 5 seconds.", name, nPanics+1)
			time.Sleep(5 * time.Second)
			go autoRestart(name, nPanics+1, f) // restart
		}
	}()
	f()
}

// PricingStream starts a stream of prices
func (streamApi *StreamAPI) PricingStream(instruments []string, pchan chan models.ClientPrice, hchan chan models.PricingHeartbeat) {

	url := streamApi.context.StreamApiURL + "/v3/accounts/" + streamApi.context.Account + "/pricing/stream"
	qurl := url + "?instruments=" + strings.Join(instruments, ",")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", qurl, nil)
	req.Header.Add("Authorization", "Bearer "+streamApi.context.Token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		reader := bufio.NewReader(response.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				log.Println(err)
				panic("Connection on clientStream is lost")
			}
			var p models.ClientPrice
			json.Unmarshal([]byte(line), &p)
			if p.Type == "HEARTBEAT" {
				h := models.PricingHeartbeat{
					Type: p.Type,
					Time: p.Time,
				}
				hchan <- h
			} else {
				pchan <- p
			}
		}
	}
	//return nil
}
