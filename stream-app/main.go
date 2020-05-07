package main

import (
	"fmt"
	"os"

	"github.com/xtordoir/goanda/api"
	"github.com/xtordoir/goanda/models"
)

func priceProcessor(c chan models.ClientPrice) {
	for {
		data := <-c
		tick := models.ClientPrice2Tick(&data)
		fmt.Println(tick)
	}
}

func heartbeatProcessor(c chan models.PricingHeartbeat) {
	for {
		data := <-c
		fmt.Printf("%s\n", data)
	}
}

func main() {

	// channels for data
	pchan := make(chan models.ClientPrice)
	hchan := make(chan models.PricingHeartbeat)

	// start processors for data
	go priceProcessor(pchan)
	go heartbeatProcessor(hchan)

	// context to create api
	ctx := api.Context{
		ApiURL:      os.Getenv("OANDA_STREAM_URL"),
		Token:       os.Getenv("OANDA_API_KEY"),
		Account:     os.Getenv("OANDA_ACCOUNT"),
		Application: "TestStreaming",
	}
	streamapi := ctx.CreateStreamAPI()
	streamapi.PricingStream([]string{"EUR_USD", "BCO_USD", "SPX500_USD"}, pchan, hchan)
}
