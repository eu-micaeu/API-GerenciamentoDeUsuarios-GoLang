package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/eu-micaeu/API-GerenciamentoDeUsuarios-GoLang/handlers"
)

func UsuarioRoutes(r *gin.Engine, db *sql.DB) {
	
	userHandler := handlers.Usuario{}

	r.POST("/login", userHandler.Entrar(db))

	r.POST("/register", userHandler.Registrar(db))

	r.POST("/perfil-token", userHandler.PegarInformacoesDoUsuarioAtravesDoToken(db))

	r.POST("/exit", userHandler.Sair(db))

}
