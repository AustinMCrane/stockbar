package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/johnmccabe/go-bitbar"
	"github.com/timpalpant/go-iex"
)

var symbolsStr = flag.String("symbols", "AAPL,FB", "symbols to show in bit bar")

func main() {
	flag.Parse()

	client := iex.NewClient(&http.Client{
		Timeout: 5 * time.Second,
	})

	symbols := strings.Split(*symbolsStr, ",")

	quotes, err := client.GetStockQuotes(symbols)
	if err != nil {
		panic(err)
	}

	app := bitbar.New()
	for i := range symbols {
		symbol := symbols[i]
		quote := quotes[symbol]

		status := fmt.Sprintf("%s: %.2f%%", symbol, quote.ChangePercent*100)
		line := app.StatusLine(status)
		if quote.Change < 0 {
			line.Color("red")
		} else {
			line.Color("green")
		}
	}
	app.Render()
}
