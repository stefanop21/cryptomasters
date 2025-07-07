package main

import (
	"fmt"
	"sync"

	"github.com/stefanop21/cryptomasters/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "BCH"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getRate(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func getRate(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
}
