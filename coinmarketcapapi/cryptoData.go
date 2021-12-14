package coinmarketcapapi

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

var apikeys, apikeyName, apikeysAddress string

func init() {
	apikeys = os.Getenv("COINMARKETCAP_APIKEYS")
	apikeyName = os.Getenv("COINMARKETCAP_APIKEYS_NAME")
	apikeysAddress = os.Getenv("COINMARKETCAP_APIKEYS_ADDRESS")
}

// Quotes Latest is https://coinmarketcap.com/api/documentation/v1/#operation/getV1CryptocurrencyQuotesLatest
type QuotesLatest struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		Btc QuotesLatestData `json:"btc"`
		Eth QuotesLatestData `json:"eth"`
	} `json:"data"`
}

type QuotesLatestData struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	Slug              string      `json:"slug"`
	NumMarketPairs    int         `json:"num_market_pairs"`
	DateAdded         time.Time   `json:"date_added"`
	Tags              []string    `json:"tags"`
	MaxSupply         int         `json:"max_supply"`
	CirculatingSupply int         `json:"circulating_supply"`
	TotalSupply       int         `json:"total_supply"`
	IsActive          int         `json:"is_active"`
	Platform          interface{} `json:"platform"`
	CmcRank           int         `json:"cmc_rank"`
	IsFiat            int         `json:"is_fiat"`
	LastUpdated       time.Time   `json:"last_updated"`
	Quote             struct {
		Usd struct {
			Price                 float64   `json:"price"`
			Volume24H             float64   `json:"volume_24h"`
			VolumeChange24H       float64   `json:"volume_change_24h"`
			PercentChange1H       float64   `json:"percent_change_1h"`
			PercentChange24H      float64   `json:"percent_change_24h"`
			PercentChange7D       float64   `json:"percent_change_7d"`
			PercentChange30D      float64   `json:"percent_change_30d"`
			PercentChange60D      float64   `json:"percent_change_60d"`
			PercentChange90D      float64   `json:"percent_change_90d"`
			MarketCap             float64   `json:"market_cap"`
			MarketCapDominance    float64   `json:"market_cap_dominance"`
			FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
			LastUpdated           time.Time `json:"last_updated"`
		} `json:"USD"`
	} `json:"quote"`
}

func GetcryptoDataList() {

	client := &http.Client{}

	req, err := http.NewRequest("Get", apikeysAddress, nil)

	if err != nil {
		fmt.Println("new api request fail")
	}

	q := url.Values{}
	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add(apikeyName, apikeys)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("coinMarketCap api request fail")
	}
	fmt.Println(resp.Body)

}
