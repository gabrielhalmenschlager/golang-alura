package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Conectar() {
	url := "host=localhost user=root password=root dbname=gogames port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados:", err)
	}

	log.Println("Banco de dados conectado com sucesso!")
}
