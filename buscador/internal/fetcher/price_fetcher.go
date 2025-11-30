package fetcher

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/models"
)

func FetchPrices(priceChannel chan<- models.PriceDetail) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- FetchPriceFromSite3()
	}()

	go func() {
		defer wg.Done()
		FetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func FetchPriceFromSite1() models.PriceDetail {
	time.Sleep(1 * time.Second)
	return models.PriceDetail{
		StoreName: "Loja A",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite2() models.PriceDetail {
	time.Sleep(2 * time.Second)
	return models.PriceDetail{
		StoreName: "Loja B",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchPriceFromSite3() models.PriceDetail {
	time.Sleep(3 * time.Second)
	return models.PriceDetail{
		StoreName: "Loja C",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func FetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}
	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "Loja D",
			Value:     price,
			Timestamp: time.Now(),
		}
	}
}
