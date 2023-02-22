package main

import "net/http"

type Watcher struct {
	Tickers  []string
	Client   http.Client
	Messages chan string
}

type ValidData struct {
	Data struct {
		Symbol         string `json:"symbol"`
		CompanyName    string `json:"companyName"`
		StockType      string `json:"stockType"`
		Exchange       string `json:"exchange"`
		IsNasdaqListed bool   `json:"isNasdaqListed"`
		IsNasdaq100    bool   `json:"isNasdaq100"`
		IsHeld         bool   `json:"isHeld"`
		PrimaryData    struct {
			LastSalePrice      string `json:"lastSalePrice"`
			NetChange          string `json:"netChange"`
			PercentageChange   string `json:"percentageChange"`
			DeltaIndicator     string `json:"deltaIndicator"`
			LastTradeTimestamp string `json:"lastTradeTimestamp"`
			IsRealTime         bool   `json:"isRealTime"`
			BidPrice           string `json:"bidPrice"`
			AskPrice           string `json:"askPrice"`
			Volume             string `json:"volume"`
		} `json:"primaryData"`
		SecondaryData any    `json:"secondaryData"`
		MarketStatus  string `json:"marketStatus"`
		AssetClass    string `json:"assetClass"`
		KeyStats      any    `json:"keyStats"`
		Notifications []any  `json:"notifications"`
	} `json:"data"`
	Message any `json:"message"`
	Status  struct {
		RCode            int `json:"rCode"`
		BCodeMessage     any `json:"bCodeMessage"`
		DeveloperMessage any `json:"developerMessage"`
	} `json:"status"`
}
