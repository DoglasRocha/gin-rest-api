package main

import (
	"gin-rest-api/models"
	"gin-rest-api/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{"Doglas", "00000000000", "000000000"},
		{"Ana", "11111111111", "111111111"},
	}
	routes.HandleRequests()
}
