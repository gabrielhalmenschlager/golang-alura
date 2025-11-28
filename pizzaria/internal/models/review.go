package models

type Review struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
