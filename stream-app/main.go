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
		ApiURL:       os.Getenv("OANDA_API_URL"),
		StreamApiURL: os.Getenv("OANDA_STREAM_URL"),
		Token:        os.Getenv("OANDA_API_KEY"),
		Account:      os.Getenv("OANDA_ACCOUNT"),
		Application:  "TestStreaming",
	}

	fmt.Printf("%s\n", ctx.ApiURL)
	fmt.Printf("%s\n", ctx.StreamApiURL)
	// fmt.Printf("%s\n", ctx.Token)
	fmt.Printf("%s\n", ctx.Account)
	fmt.Printf("%s\n", ctx.Application)

	api := ctx.CreateAPI()

	if len(ctx.Account) == 0 {
		accounts, err := api.GetAccounts()
		if err == nil && len(accounts.Accounts) > 0 {
			fmt.Printf("Setting Account # in context: %s\n", accounts.Accounts[0].ID)
			ctx.Account = accounts.Accounts[0].ID
		}
	}

	pos, err := api.GetPricing([]string{"EUR_USD"})

	// err := api.GetAccounts()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	fmt.Printf("%s\n", pos)

	streamapi := ctx.CreateStreamAPI()
	streamapi.PricingStream([]string{"EUR_USD", "BCO_USD", "SPX500_USD", "EUR_JPY"}, pchan, hchan)
	// streamapi.PricingStream([]string{"EUR_USD"}, pchan, hchan)
}
