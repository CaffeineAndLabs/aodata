package aodata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is the HTTP Client used to contact albion-online-data
type Client struct {
	BaseURL    string
	HttpClient *http.Client
}

// Price reprensents the price of an item in the albion-online-data API
type Price struct {
	ItemID           string `json:"item_id,omitempty"`
	City             string `json:"city,omitempty"`
	Quality          int    `json:"quality,omitempty"`
	SellPriceMin     int    `json:"sell_price_min,omitempty"`
	SellPriceMinDate string `json:"sell_price_min_date,omitempty"`
	SellPriceMax     int    `json:"sell_price_max,omitempty"`
	SellPriceMaxDate string `json:"sell_price_max_date,omitempty"`
	BuyPriceMin      int    `json:"buy_price_min,omitempty"`
	BuyPriceMinDate  string `json:"buy_price_min_date,omitempty"`
	BuyPriceMax      int    `json:"buy_price_max,omitempty"`
	BuyPriceMaxDate  string `json:"buy_price_max_date,omitempty"`
}

func NewClient() *Client {
	return &Client{
		BaseURL:    "https://www.albion-online-data.com/api/v2",
		HttpClient: &http.Client{},
	}
}

// GetPrices gets all prices for a specific items (in multiple regions)
func (c *Client) GetPrices(itemID string) ([]Price, error) {
	url := fmt.Sprintf(c.BaseURL+"/stats/Prices/%s", itemID)
	resp, err := c.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var prices []Price
	json.Unmarshal(body, &prices)

	return prices, nil
}
