package api

import (
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
	fmt.Println(string(data))
	positions, _ := parseAccountOpenPositions(&data)
	fmt.Println(positions)

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
