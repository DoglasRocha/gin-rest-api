package main

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func SetupDasRotasDeTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestFalhador(t *testing.T) {
	t.Fatalf("teste falhou")
}
