package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/models"
	"go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{}

	database.ConectaComBancoDeDados()
	fmt.Println("Conectado ao banco")
	routes.HandleRequest()
}
