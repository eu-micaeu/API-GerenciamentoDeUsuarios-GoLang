package handlers

import (
	"database/sql"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	User_ID int `json:"id_usuario"`
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

// Função para verificar se o token está na tabela de tokens inválidos
func tokenEstaNaTabelaDeTokensInvalidos(db *sql.DB, tokenString string) bool {

	query := "SELECT COUNT(*) FROM tokens_invalidos WHERE token_invalido = $1"

	var count int

	err := db.QueryRow(query, tokenString).Scan(&count)

	if err != nil {

		log.Println("Erro ao consultar a tabela de tokens inválidos:", err)

		return true

	}

	return count > 0

}

// Função com finalidade de validação do token para as funções do usuário.
func (u *User) ValidarOToken(db *sql.DB, tokenString string) (int, error) {

	if tokenEstaNaTabelaDeTokensInvalidos(db, tokenString) {

		return 0, nil

	}

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil

	})

	if err != nil || !tkn.Valid {

		return 0, err

	}

	return claims.User_ID, nil

}

// Função com finalidade de geração do token para as funções do usuário.
func GerarOToken(usuario User) (string, error) {

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{

		User_ID: usuario.User_ID,

		StandardClaims: jwt.StandardClaims{

			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {

		return "Erro ao gerar token", err

	}

	return tokenString, nil
}