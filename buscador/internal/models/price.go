package models

import "time"

type PriceDetail struct {
	StoreName string
	Value     float64
	Timestamp time.Time
}
