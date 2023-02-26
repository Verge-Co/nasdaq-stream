# nasdaq-stream
Easy and simple package to stream real-time stock prices from the NASDAQ api.

## Usage
```go
package main

import (
	"fmt"
	"net/http"
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
```
