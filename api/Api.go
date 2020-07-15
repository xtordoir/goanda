package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/xtordoir/goanda/models"
)

// API is an api instance with a context to call endpoints
type API struct {
	context Context
}

// GetOpenPositions gets the open Positions on the account
func (api *API) GetOpenPositions() (*models.AccountPositions, error) {
	client := &http.Client{}
	account := api.context.Account
	apiURL := api.context.ApiURL
	token := api.context.Token
	req, errr := http.NewRequest("GET", apiURL+"/v3/accounts/"+account+"/openPositions", nil)
	if errr != nil {
		return nil, errr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	positions, _ := parseAccountOpenPositions(&data)
	//fmt.Println(positions)

	return &positions, nil
}

// GetPosition gets the Position on the account
func (api *API) GetPosition(instrument string) (*models.AccountPosition, error) {
	client := &http.Client{}
	account := api.context.Account
	apiURL := api.context.ApiURL
	token := api.context.Token
	req, errr := http.NewRequest("GET", apiURL+"/v3/accounts/"+account+"/positions/"+instrument, nil)
	if errr != nil {
		return nil, errr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	positions, _ := parseAccountPosition(&data)
	//fmt.Println(positions)

	return &positions, nil
}

// GetPricing fetches the prricing for a list of instruments
func (api *API) GetPricing(instruments []string) (*models.Prices, error) {

	instrumentsQstr := strings.Join(instruments, ",")
	// TODO DEDUPLICATE THIS
	client := &http.Client{}
	account := api.context.Account
	apiURL := api.context.ApiURL
	token := api.context.Token
	req, errr := http.NewRequest("GET", apiURL+"/v3/accounts/"+account+"/pricing?instruments="+instrumentsQstr, nil)
	if errr != nil {
		return nil, errr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	prices, _ := parsePrices(&data)
	//fmt.Println(positions)

	return &prices, nil
}

// GetCandles fetches a number of candles for a given instrument and granularity
func (api *API) GetCandles(instrument string, num int, granularity string) (*models.Candles, error) {
	// TODO DEDUPLICATE THIS
	client := &http.Client{}
	account := api.context.Account
	apiURL := api.context.ApiURL
	token := api.context.Token
	qStr := fmt.Sprintf("?granularity=%s&count=%d", granularity, num)
	req, errr := http.NewRequest("GET", apiURL+"/v3/accounts/"+account+"/instruments/"+instrument+"/candles"+qStr, nil)
	if errr != nil {
		return nil, errr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	candles, _ := parseCandles(&data)
	//fmt.Println(positions)

	return &candles, nil
}

// PostMarketOrder posts a Market orderr a number of candles for a given instrument and granularity
func (api *API) PostMarketOrder(instrument string, units int64) (error, error) {

	orderReq := models.OrderRequest{
		Order: models.MakeMarketOrder(instrument, units),
	}
	payload, _ := json.Marshal(orderReq)

	// TODO DEDUPLICATE THIS
	client := &http.Client{}
	account := api.context.Account
	apiURL := api.context.ApiURL
	token := api.context.Token
	req, errr := http.NewRequest("POST", apiURL+"/v3/accounts/"+account+"/orders",
		bytes.NewBuffer(payload))
	if errr != nil {
		return nil, errr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	//	orderStatus, _ := parseOrderStatus(&data)
	//fmt.Println(positions)

	return nil, nil
}

// GetPositionBook fetches the last PositionBook for instruments
func (api *API) GetPositionBook(instrument string) (*models.PositionBook, error) {

	// TODO DEDUPLICATE THIS
	client := &http.Client{}
	apiURL := api.context.ApiURL
	token := api.context.Token
	req, errr := http.NewRequest("GET", apiURL+"/v3/instruments/"+instrument+"/positionBook", nil)
	if errr != nil {
		return nil, errr
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}
	data, _ := ioutil.ReadAll(response.Body)
	//fmt.Println(string(data))
	positionBook, _ := parsePositionBook(&data)
	//fmt.Println(positions)

	return &positionBook, nil
}
