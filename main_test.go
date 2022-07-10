package main

import (
	"gin-rest-api/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupDasRotasDeTest() *gin.Engine {
	rotas := gin.Default()
	return rotas
}

func TestVerificaStatusCodeDaSaudacao(t *testing.T) {
	r := SetupDasRotasDeTest()
	r.GET("/:nome", controllers.Saudacao)
	req, _ := http.NewRequest("GET", "/doglas", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Os Status Code deveriam ser iguais")
	mockDaResp := `{"API diz:":"E ai doglas, tudo bem?"}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockDaResp, string(respBody))
}
