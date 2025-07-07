package main

import (
	"fmt"

	"github.com/stefanop21/cryptomasters/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)
}
