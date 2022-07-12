package controllers

import (
	"gin-rest-api/database"
	"gin-rest-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExibeTodosAlunos godoc
// @Summary Exibe todos os alunos
// @Description Rota para exibir todos os alunos
// @Tags alunos
// @Produce json
// @Success 200 {object} models.Aluno true "Modelo do aluno"
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// Saudacao godoc
// @Summary API faz uma saudação
// @Description Rota que recebendo um nome, saúda o usuário com o nome fornecido
// @Tags saudacao
// @Produce json
// @Success 200 {object} gin.H
// @Router /:nome [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo bem?",
	})
}

// CriaNovoAluno godoc
// @Summary Cria um novo aluno
// @Description Rota para criar um novo aluno
// @Tags aluno
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "Modelo do aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H
// @Router /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorId godoc
// @Summary Busca um aluno por id
// @Description Rota para buscar um aluno por id
// @Tags aluno
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H
// @Router /alunos/:id [get]
func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
// @Summary Deleta um aluno
// @Description Rota para deletar um aluno, por id
// @Tags aluno
// @Produce json
// @Success 200 {object} models.Aluno
// @Router /alunos/:id [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAluno godoc
// @Summary Edita um aluno
// @Description Rota para editar um aluno, por id
// @Tags aluno
// @Accept json
// @Produce json
// @@Param aluno body models.Aluno true "Modelo do aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H
// @Router /alunos/:id [patch]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := models.ValidaDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
// @Summary Busca um aluno por cpf
// @Description Rota para buscar um aluno por cpf
// @Tags aluno
// @Produce json
// @Success 200 {object} models.Aluno
// @Failure 400 {object} gin.H
// @Router /alunos/cpf/:cpf [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
