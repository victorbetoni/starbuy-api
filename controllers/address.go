package controllers

import (
	"database/sql"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAddresses(c *gin.Context) error {
	queried := c.Param("user")
	user, _ := authorization.ExtractUser(c)

	if user != queried {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "authorized"})
		return nil
	}

	var addresses []model.RawAddress
	if err := repository.DownloadAddresses(user, &addresses); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, addresses)
	return nil
}

func GetAddress(c *gin.Context) error {
	queried, id := c.Param("user"), c.Param("id")

	user, _ := authorization.ExtractUser(c)

	if user != queried {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		return nil
	}

	var address model.Address
	if err := repository.DownloadAddress(id, &address); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, address)
	return nil
}

func PostAddress(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)

	type Request struct {
		CEP        string `json:"cep"`
		Number     int    `json:"number"`
		Complement string `json:"complement"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	//TODO: Usar alguma API para verificar se o CEP bate com algum existente

	address := model.RawAddress{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		Holder:     user,
		CEP:        req.CEP,
		Number:     req.Number,
		Complement: req.Complement,
	}

	c.JSON(http.StatusOK, address)
	return nil
}
