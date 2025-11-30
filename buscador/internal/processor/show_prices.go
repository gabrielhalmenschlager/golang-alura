package processor

import (
	"fmt"

	"github.com/gabrielhalmenschlager/curso-golang-alura/buscador/internal/models"
)

func ShowPriceAVG(priceChannel <-chan models.PriceDetail, done chan<- bool) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price.Value
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("[%s] Preço recebido da %s | R$ %.2f | Preço médio até agora: R$ %.2f \n", price.Timestamp.Format("02-Jan 15:04:05"), price.StoreName, price.Value, avgPrice)
	}
	done <- true
}
