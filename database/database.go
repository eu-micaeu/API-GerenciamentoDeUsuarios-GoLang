package database

import (

	"database/sql"

	"fmt"
	"log"

	_ "github.com/lib/pq"

)

func NewDB() (*sql.DB, error) {

	dbUser := "postgres"

	dbPassword := "12345678"

	dbHost := "localhost"

	dbPort := "5432"

	dbName := "postgres"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)

	if err != nil {

		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)

	}

	log.Println("Conectado ao banco de dados com sucesso!")

	return db, nil
}
