package main

import (
	
	"github.com/gin-gonic/gin"

	"github.com/eu-micaeu/API-GerenciamentoDeUsuarios-GoLang/database"

	"github.com/eu-micaeu/API-GerenciamentoDeUsuarios-GoLang/middlewares"

	"github.com/eu-micaeu/API-GerenciamentoDeUsuarios-GoLang/routes"

)

func main() {

	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	db, err := database.NewDB()

	if err != nil {

		panic(err)

	}

	routes.UserRoutes(r, db)

	r.Run()
}
