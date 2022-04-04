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
	Username string `db:"username"`
	Password string `db:"password"`
}

func Auth(c *gin.Context) {
	db := database.GrabDB()
	login := Login{}

	if err := c.BindJSON(&login); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	recorded := Login{}
	if err := db.Get(&recorded, "SELECT * FROM login WHERE username=$1", login.Username); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, errors.New("not found"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user model.User
	if err := repository.DownloadUser(login.Username, &user); err != nil {
		c.AbortWithError(http.StatusNotFound, errors.New("user not found"))
		return
	}

	if err := security.ComparePassword(recorded.Password, login.Password); err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	token := authorization.GenerateToken(login.Username)

	type Response struct {
		User  model.User `json:"user"`
		Token string     `json:"jwt"`
	}
	c.JSON(http.StatusOK, Response{User: user, Token: token})
}
