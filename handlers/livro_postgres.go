package handlers

import (
	"crud-biblioteca/repository/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetLivrosPostgresHandler(repo *postgres.LivroRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		livros, err := repo.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livros"})
			return
		}
		c.JSON(http.StatusOK, livros)
	}
}

// Implemente outros handlers (POST, PUT, DELETE) conforme necessário
