package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
	"trader/market"
)

func main() {
	markets := make(market.Markets)
	var trade market.Trade
	beginTime := time.Now()

	// read data from standin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
			line := scanner.Text()

			err := json.Unmarshal([]byte(line), &trade)
			if (err != nil) {
				continue
			}

			if markets[trade.Market] == nil {
				markets[trade.Market], _ = market.NewMarket()
			}

			markets[trade.Market].UpdateMarket(&trade)
	}


	for _, market := range markets {
		marketJson, err := json.Marshal(market)
		if (err != nil) {
			continue
		}
		fmt.Println(string(marketJson))
	}

	fmt.Println(fmt.Sprintf("Duration of send operation: %s", time.Now().Sub(beginTime).String()))
}
