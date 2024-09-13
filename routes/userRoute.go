package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/eu-micaeu/API-GerenciamentoDeUsuarios-GoLang/handlers"
)

func UserRoutes(r *gin.Engine, db *sql.DB) {
	
	userHandler := handlers.User{}

	r.GET("/login", userHandler.Entrar(db))

	r.POST("/register", userHandler.Registrar(db))

	r.PUT("/update", userHandler.Atualizar(db))

	r.DELETE("/delete", userHandler.Deletar(db))

	r.POST("/exit", userHandler.Sair(db))

}
