package database

import (

	"database/sql"

	"fmt"
	"log"

	_ "github.com/lib/pq"

)

func NewDB() (*sql.DB, error) {

	dbUser := "root"

	dbPassword := "c9N1ydwtNIlvwjl8LxHQPCgQ3xmXrcvg"

	dbHost := "dpg-cmgs7b6n7f5s739moo2g-a.oregon-postgres.render.com"

	dbPort := "5432"

	dbName := "vocatito"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)

	if err != nil {

		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %v", err)

	}

	if err = db.Ping(); err != nil {

		return nil, fmt.Errorf("erro ao pingar o banco de dados: %v", err)
		
	}

	log.Println("Conectado ao banco de dados com sucesso!")

	return db, nil
}
