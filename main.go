package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {
	tickers := []string{"MSFT", "AAPL"}

	var watcher = Watcher{
		Tickers:  tickers,
		Client:   http.Client{},
		Messages: make(chan string),
	}

	for _, ticker := range tickers {
		go watcher.watchTicker(ticker)
	}

	for {
		msg := <-watcher.Messages

		fmt.Println(msg)
	}
}

func (w Watcher) watchTicker(ticker string) error {
	url := fmt.Sprintf("https://api.nasdaq.com/api/quote/%s/info?assetclass=stocks", ticker)

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Verge-Watcher/0.0.1")

	for {
		res, err := w.Client.Do(req)

		if err != nil {
			return err
		}

		if res.StatusCode != 200 {
			return errors.New(fmt.Sprintf("invalid status, %v", res.StatusCode))
		}

		var data ValidData

		err = json.NewDecoder(res.Body).Decode(&data)

		if err != nil {
			return err
		}

		price := data.Data.PrimaryData.LastSalePrice[1:]

		w.Messages <- fmt.Sprintf("%s: %s", ticker, price)

		time.Sleep(time.Second)
	}
}
