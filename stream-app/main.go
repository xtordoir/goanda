package streamapp

import (
	"fmt"
	"os"

	"github.com/xtordoir/goanda/api"
	"github.com/xtordoir/goanda/models"
)

func main() {
	ctx := api.Context{
		ApiURL:      os.Getenv("OANDA_STREAM_URL"),
		Token:       os.Getenv("OANDA_API_KEY"),
		Account:     os.Getenv("OANDA_ACCOUNT"),
		Application: "TestStreaming",
	}
	streamapi := ctx.CreateStreamAPI()
	streamapi.PricingStream([]string{"EUR_USD"}, func(p *models.ClientPrice) { fmt.Println(*p) }, func(p *models.PricingHeartbeat) { fmt.Println(*p) })
}
