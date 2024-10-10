package main

import "fmt"

func exchangeToUSD(amountIDR float64, exchangeRateUSD_IDR float64, amountSGD float64, exchangeRateUSD_SGD float64) float64 {
	idrToUSD := amountIDR / exchangeRateUSD_IDR
	sgdToUSD := amountSGD / exchangeRateUSD_SGD
	totalUSD := idrToUSD + sgdToUSD
	return totalUSD
}

func main() {
	amountIDR := 120000.0
	exchangeRateUSD_IDR := 15000.0
	amountSGD := 29.0
	exchangeRateUSD_SGD := 1.5

	totalUSD := exchangeToUSD(amountIDR, exchangeRateUSD_IDR, amountSGD, exchangeRateUSD_SGD)
	fmt.Printf("Total USD: %.6f\n", totalUSD)
}
