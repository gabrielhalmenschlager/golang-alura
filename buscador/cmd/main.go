package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/fetcher"
)

func main() {
	start := time.Now()
	var price1, price2, price3 float64
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		price1 = fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		price2 = fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		price3 = fetcher.FetchPriceFromSite3()
	}()

	wg.Wait()

	fmt.Printf("R$ %.2f \n", price1)
	fmt.Printf("R$ %.2f \n", price2)
	fmt.Printf("R$ %.2f \n", price3)

	fmt.Printf("\nTempo total: %s", time.Since(start))
}
