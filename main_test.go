package main

import (
	"bytes"
	"encoding/json"
	"gin-rest-api/controllers"
	"gin-rest-api/database"
	"gin-rest-api/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "aluno teste", CPF: "12345678910", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeDaSaudacao(t *testing.T) {
	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/doglas", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Os Status Code deveriam ser iguais")
	mockDaResp := `{"API diz:":"E ai doglas, tudo bem?"}`
	respBody, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, mockDaResp, string(respBody))
}

func TestListaTodosOsAlunosHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678910", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorId)

	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", pathDaBusca, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var alunoMock models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMock)

	assert.Equal(t, "aluno teste", alunoMock.Nome)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, "12345678910", alunoMock.CPF)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestDeletaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	pathDeBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathDeBusca, nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestEditaAlunoHandler(t *testing.T) {
	database.ConectaComBancoDeDados()

	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "aluno teste", CPF: "47123456789", RG: "123456780"}
	alunoJson, _ := json.Marshal(aluno)

	req, _ := http.NewRequest("PATCH", "/alunos/"+strconv.Itoa(ID), bytes.NewBuffer(alunoJson))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	var alunoMockAtualizado models.Aluno
	json.Unmarshal(resp.Body.Bytes(), &alunoMockAtualizado)
	assert.Equal(t, "47123456789", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456780", alunoMockAtualizado.RG)
	assert.Equal(t, "aluno teste", alunoMockAtualizado.Nome)
}
