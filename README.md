# goanda
Golang oanda REST-V20 API library

See https://developer.oanda.com/rest-live-v20/introduction/ for documentation.

## Quick start



Create an api context:

```
ctx := api.Context{
  ApiURL:       apiURL,
  StreamApiURL: streamURL,
  Token:        token,
  Account:      account,
  Application:  "MyApp",
}
```

Use the configured context to create an api instance and make calls:

```
api := ctx.CreateAPI()
pos, err := api.GetPricing([]string{"EUR_USD", "EUR_JPY"})
```

Or use the configured context to create a streaming api instance, and open a pricing stream. Prices are sent to a Price channel and heartbeat to another:

```
streamapi := ctx.CreateStreamAPI()

pchan := make(chan models.ClientPrice)
hchan := make(chan models.PricingHeartbeat)

go streamapi.PricingStream([]string{"EUR_USD", "SPX500_USD"}, pchan, hchan)
```


## API Endpoints

Implemented Endpoints are in the `api` sub-package (Api.go):

- **GetOpenPositions**: Read Open positions for the account:
```
func (api *API) GetOpenPositions() (*models.AccountPositions, error)
```

- **GetPosition**: Read a specific instrument position for the account :

```
func (api *API) GetPosition(instrument string) (*models.AccountPosition, error)
```

- **GetPricing**: Get the current price for a list of instruments:

```
func (api *API) GetPricing(instruments []string) (*models.Prices, error)
```

- **GetCandles**: Get `num` historical candles (OHLC) for an instrument, and granularity (see oanda definitions):

```
func (api *API) GetCandles(instrument string, num int, granularity string) (*models.Candles, error)
```

- **PostMarketOrder**: Send a market order for an instrument, it is a synchronous call. Note: returned data is not implemented, it is an optimistic implementation!!!

```
func (api *API) PostMarketOrder(instrument string, units int64) (error, error)
```

- **GetPositionBook**: Get the aggregate positions from oanda customers:

```
func (api *API) GetPositionBook(instrument string) (*models.PositionBook, error)
```

- **GetAccounts()**: Get the list of accounts for the api key:

```
func (api *API) GetAccounts() (*models.Accounts, error)
```


## Streaming API Endpoints

Implemented Endpoints for Streaming are in the `api` sub-package (StreamApi.go):

- **PricingStream**: Used to start a stream for selected instruments. Prices and Heartbeats are sent to channels.

```
func (streamApi *StreamAPI) PricingStream(instruments []string, pchan chan models.ClientPrice, hchan chan models.PricingHeartbeat)
```
## Oanda Definitions

TODO: Complete implemented definition list, see models sub-package for up-to-date information

Implemented Definitions are in the `models` sub-package:

- **Position Definitions**
  - Position
  - PositionSide
- **Pricing Definitions**
  - ClientPrice
  - PricingHeartbeat
- **Pricing Common Definitions**
  - PriceBucket
- **Primitives Definitions**

## Extra Definitions

- **Tick**
