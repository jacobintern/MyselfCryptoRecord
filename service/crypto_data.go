package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// 改由 db 取得
var favCryptoList = []string{
	"BTC",
	"ETH",
	"GNX",
	"FIL",
	"ADA",
	"MANA",
	"LINK",
	"KNC",
	"FTT",
	"SHIB",
}

// 購買的價格，改從資料庫取得
var purchasePrice = map[string]float64{
	"GNX":  0.04723,
	"FIL":  87,
	"ADA":  2.543,
	"MANA": 3.32,
	"LINK": 30.54,
	"KNC":  1.67,
	"FTT":  67.458,
	"SHIB": 0.0000361,
}

// 擁有的數量，改從資料庫取得
var myWallet = map[string]float64{
	"GNX":  11014,
	"FIL":  5.7,
	"ADA":  168.08,
	"MANA": 30.06,
	"LINK": 5.3,
	"KNC":  59.67,
	"FTT":  0.734,
	"SHIB": 1969425.08,
}

type ApiInfoModel struct {
	apikeys    string
	apikeyName string
	quotesApi  string
	mapApi     string
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
		Btc  QuotesLatestData `json:"btc"`
		Eth  QuotesLatestData `json:"eth"`
		Ftt  QuotesLatestData `json:"ftt"`
		Shib QuotesLatestData `json:"shib"`
		Gnx  QuotesLatestData `json:"gnx"`
		Fil  QuotesLatestData `json:"fil"`
		Ada  QuotesLatestData `json:"ada"`
		Mana QuotesLatestData `json:"mana"`
		Link QuotesLatestData `json:"link"`
		Knc  QuotesLatestData `json:"knc"`
	} `json:"data"`
}

type QuotesLatestData struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Slug   string `json:"slug"`
	// NumMarketPairs int       `json:"num_market_pairs"`
	// DateAdded      time.Time `json:"date_added"`
	// Tags              []string    `json:"tags"`
	// MaxSupply         int         `json:"max_supply"`
	// CirculatingSupply float64     `json:"circulating_supply"`
	// TotalSupply       float64     `json:"total_supply"`
	IsActive int `json:"is_active"`
	// Platform          interface{} `json:"platform"`
	CmcRank int `json:"cmc_rank"`
	// IsFiat            int         `json:"is_fiat"`
	LastUpdated time.Time `json:"last_updated"`
	Quote       struct {
		Usd struct {
			Price float64 `json:"price"`
			// Volume24H             float64   `json:"volume_24h"`
			// VolumeChange24H       float64   `json:"volume_change_24h"`
			// PercentChange1H       float64   `json:"percent_change_1h"`
			// PercentChange24H      float64   `json:"percent_change_24h"`
			// PercentChange7D       float64   `json:"percent_change_7d"`
			// PercentChange30D      float64   `json:"percent_change_30d"`
			// PercentChange60D      float64   `json:"percent_change_60d"`
			// PercentChange90D      float64   `json:"percent_change_90d"`
			// MarketCap             float64   `json:"market_cap"`
			// MarketCapDominance    float64   `json:"market_cap_dominance"`
			// FullyDilutedMarketCap float64   `json:"fully_diluted_market_cap"`
			LastUpdated time.Time `json:"last_updated"`
		} `json:"USD"`
	} `json:"quote"`
}

var _apiInfo ApiInfoModel

func init() {
	_apiInfo = ApiInfoModel{
		apikeys:    os.Getenv("COINMARKETCAP_APIKEYS"),
		apikeyName: os.Getenv("COINMARKETCAP_APIKEYS_NAME"),
		quotesApi:  os.Getenv("COINMARKETCAP_QUOTES_LASTEST"),
		mapApi:     os.Getenv("COINMARKETCAP_MAP"),
	}
}

func GetcryptoDataList() QuotesLatest {

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, _apiInfo.quotesApi, nil)

	if err != nil {
		fmt.Println("new api request fail")
	}

	// 增加查詢的參數
	params := req.URL.Query()
	params.Add("symbol", "BTC,ETH,FTT,SHIB,GNX,FIL,ADA,MANA,LINK,KNC")
	params.Add("convert", "USD")
	req.URL.RawQuery = params.Encode()

	// 設定 head 參數
	req.Header.Set("Accepts", "application/json")
	req.Header.Set(_apiInfo.apikeyName, _apiInfo.apikeys)

	// 發送請求
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("coinMarketCap api request fail")
	}

	// 關閉數據
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	body, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		fmt.Println(err)
	}

	// 宣告 model
	var cryptoList QuotesLatest

	// json string 轉換 model
	jsonErr := json.Unmarshal(body, &cryptoList)

	if jsonErr != nil {
		fmt.Println(jsonErr)
		return QuotesLatest{}
	}

	return cryptoList
}

func GetMapList() {

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodGet, _apiInfo.mapApi, nil)

	req.Header.Add(_apiInfo.apikeyName, _apiInfo.apikeys)

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res.Status)
	fmt.Println(string(body))
}

func MappingMyList() error {
	cryptoList := GetcryptoDataList()

	if cryptoList.Status.ErrorCode != 0 {
		err := errors.New("CoinMarketCap 資料取得取得失敗")
		return err
	}

	// 呼叫自己錢包的資訊
	for _, target := range favCryptoList {
		switch target {
		case "FTT":

		case "SHIB":

		case "GNX":

		case "FIL":

		case "ADA":

		case "MANA":

		case "LINK":

		case "KNC":

		}
	}

	return nil
}
