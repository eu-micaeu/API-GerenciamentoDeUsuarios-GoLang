package handlers

import (

	"database/sql"

	"net/http"

	"time"

	"github.com/gin-gonic/gin"

)

type User struct {

	User_ID   int       `json:"user_id"`

	Username  string    `json:"username"`

	Email     string    `json:"email"`

	Password  string    `json:"password"`

	FullName  string    `json:"full_name"`
	
	CreatedAt time.Time `json:"created_at"`
}

// Função com finalidade de login do usuário.
func (u *User) Entrar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var user User

		if err := c.BindJSON(&user); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao fazer login"})

			return

		}

		row := db.QueryRow("SELECT username, password FROM users WHERE username = $1 AND password = $2", user.Username, user.Password)

		err := row.Scan(&user.Username, &user.Password)

		if err != nil {

			c.JSON(404, gin.H{"message": "Usuário ou senha incorretos"})

			return

		}

		token, _ := GerarOToken(user)

		http.SetCookie(c.Writer, &http.Cookie{

			Name: "token",

			Value: token,

			Expires: time.Now().Add(24 * time.Hour),

			HttpOnly: true,

			Secure: true,

			SameSite: http.SameSiteStrictMode,
		})

		c.JSON(200, gin.H{"message": "Login efetuado com sucesso!", "token": token, "usuario": user})

	}

}

// Função com finalidade de registrar um usuário no sistema.
func (u *User) Registrar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var newUser User

		if err := c.BindJSON(&newUser); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao criar usuario"})

			return

		}

		_, err := db.Exec("INSERT INTO users (username, email, password, full_name, created_at) VALUES ($1, $2, $3, $4, $5, $6)", newUser.Username, newUser.Email, newUser.Password, newUser.FullName, time.Now())

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao criar usuário"})

			return

		}

		c.JSON(200, gin.H{"message": "Usuário criado com sucesso!"})

	}

}

// Função com finalidade de atualizar um usuário no sistema.
func (u *User) Atualizar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var user User

		token, err := c.Cookie("token")

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao atualizar usuario"})

			return

		}

		user_id, _ := u.ValidarOToken(token)

		if err := c.BindJSON(&user); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao atualizar usuario"})

			return

		}

		_, err = db.Exec("UPDATE users SET username = $1, email = $2, password = $3, full_name = $4 WHERE user_id = $5", user.Username, user.Email, user.Password, user.FullName, user_id)

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao atualizar usuário"})

			return

		}

		c.JSON(200, gin.H{"message": "Usuário atualizado com sucesso!"})

	}

}

// Função com finalidade de deletar um usuário no sistema.
func (u *User) Deletar(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		var user User

		token, err := c.Cookie("token")

		if err != nil {

			c.JSON(400, gin.H{"message": "Erro ao deletar usuario"})

			return

		}

		user_id, _ := u.ValidarOToken(token)

		if err := c.BindJSON(&user); err != nil {

			c.JSON(400, gin.H{"message": "Erro ao deletar usuario"})

			return

		}

		_, err = db.Exec("DELETE FROM users WHERE user_id = $1", user_id)

		if err != nil {

			c.JSON(500, gin.H{"message": "Erro ao deletar usuário"})

			return

		}

		c.JSON(200, gin.H{"message": "Usuário deletado com sucesso!"})

	}

}

// Função com finalidade de login do usuário.
func (u *User) Sair(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {

		cookie := &http.Cookie{

			Name: "token",

			Value: "",

			Expires: time.Unix(0, 0),

			HttpOnly: true,

			Secure: true,

			SameSite: http.SameSiteStrictMode,

			Path: "/",
		}

		http.SetCookie(c.Writer, cookie)

		c.JSON(200, gin.H{"message": "Saiu com sucesso!"})

	}

}
