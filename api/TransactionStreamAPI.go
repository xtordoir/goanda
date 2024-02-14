package api

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/xtordoir/goanda/models"
)

// StreamAPI is an api instance with a context to call endpoints
type TransactionStreamAPI struct {
	context Context
}

type transactionProcessor func(p *models.Transaction)
type transactionHeartbeatProcessor func(p *models.TransactionHeartbeat)

// TickStream starts a stream of transactions
func (streamApi *TransactionStreamAPI) StartTransactionStream(tchan chan models.Transaction, hchan chan models.TransactionHeartbeat) {

	// AutoRestart for TransactionStream
	go autoRestart("TransactionStream", 0, func() { streamApi.TransactionStream(tchan, hchan) })

	fmt.Println("Starting loop on Transactions")
	// for {
	// 	price := <-pchan
	// 	tchan <- models.ClientPrice2Tick(&price)
	// }
}

// transactionStreamAutoRestart for the TransactionStream function as connection reset can result in panic
func transactionStreamAutoRestart(name string, nPanics int64, f func()) {
	defer func() {
		if v := recover(); v != nil {
			// A panic is detected.
			log.Printf("%s is crashed. Panic #%d. Restarting in 5 seconds.", name, nPanics+1)
			time.Sleep(5 * time.Second)
			go transactionStreamAutoRestart(name, nPanics+1, f) // restart
		}
	}()
	f()
}

// TransactionStream starts a stream of prices
func (streamApi *TransactionStreamAPI) TransactionStream(tchan chan models.Transaction, hchan chan models.TransactionHeartbeat) {

	url := streamApi.context.StreamApiURL + "/v3/accounts/" + streamApi.context.Account + "/transactions/stream"
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
			var p models.Transaction
			json.Unmarshal([]byte(line), &p)
			if p.Type == "HEARTBEAT" {
				var h models.TransactionHeartbeat
				json.Unmarshal([]byte(line), &h)
				hchan <- h
			} else {
				tchan <- p
			}
		}
	}
	//return nil
}
