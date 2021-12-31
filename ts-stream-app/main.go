package main

import (
	"fmt"
	"os"

	"github.com/xtordoir/goanda/api"
	"github.com/xtordoir/goanda/models"
)

func transactionProcessor(c chan models.Transaction) {
	for {
		transaction := <-c
		fmt.Println(transaction)
	}
}

func heartbeatProcessor(c chan models.TransactionHeartbeat) {
	for {
		data := <-c
		fmt.Printf("%s\n", data)
	}
}

func main() {

	// channels for data
	tchan := make(chan models.Transaction)
	hchan := make(chan models.TransactionHeartbeat)

	// start processors for data
	go transactionProcessor(tchan)
	go heartbeatProcessor(hchan)

	// context to create api
	ctx := api.Context{
		ApiURL:       os.Getenv("OANDA_URL"),
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
			api = ctx.CreateAPI()
		}
	}

	pos, err := api.GetPricing([]string{"EUR_USD"})

	// err := api.GetAccounts()
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	fmt.Printf("%s\n", pos)

	streamapi := ctx.CreateTransactionStreamAPI()
	//streamapi.PricingStream([]string{"EUR_USD", "BCO_USD", "SPX500_USD", "EUR_JPY"}, pchan, hchan)

	streamapi.StartTransactionStream(tchan, hchan)
	select{}
}
