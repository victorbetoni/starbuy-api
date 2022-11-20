package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/security"

	"github.com/gin-gonic/gin"
)

type Login struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func Auth(c *gin.Context) (int, error) {
	db := database.GrabDB()
	login := Login{}

	if err := c.BindJSON(&login); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	recorded := Login{}
	if err := db.Get(&recorded, "SELECT * FROM login WHERE username=$1", login.Username); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Usuário não encontrado", "user": "", "jwt": ""})
			return 0, nil
		}
		return http.StatusInternalServerError, err
	}

	var user model.User
	if err := repository.DownloadUser(login.Username, &user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Usuário não encontrado", "user": nil, "jwt": ""})
		return 0, nil
	}

	if err := security.ComparePassword(recorded.Password, login.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Senha incorreta", "user": nil, "jwt": ""})
		return 0, nil
	}

	token := authorization.GenerateToken(login.Username)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Sessão iniciada com sucesso", "user": user, "jwt": token})

	return 0, nil
}
