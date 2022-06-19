package controllers

import (
	"database/sql"
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

func Auth(c *gin.Context) error {
	db := database.GrabDB()
	login := Login{}

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request", "user": "", "jwt": ""})
		return nil
	}

	recorded := Login{}
	if err := db.Get(&recorded, "SELECT * FROM login WHERE username=$1", login.Username); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found", "user": "", "jwt": ""})
			return nil
		}
		return err
	}

	var user model.User
	if err := repository.DownloadUser(login.Username, &user); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found", "user": nil, "jwt": ""})
		return nil
	}

	if err := security.ComparePassword(recorded.Password, login.Password); err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "Senha incorreta", "user": nil, "jwt": ""})
		return nil
	}

	token := authorization.GenerateToken(login.Username)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Sessão iniciada com sucesso", "user": user, "jwt": token})

	return nil
}
