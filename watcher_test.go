package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWatcher(t *testing.T) {
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
