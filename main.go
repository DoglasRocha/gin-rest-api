package main

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"gin-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Doglas", CPF: "00000000000", RG: "000000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "111111111"},
	}
	routes.HandleRequests()
}
