package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAddresses(c *gin.Context) {
	queried := c.Param("user")
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	if user != queried {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	var addresses []model.RawAddress
	if err := repository.DownloadAddresses(user, &addresses); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNoContent, errors.New("no content"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, addresses)
}

func GetAddress(c *gin.Context) {
	queried, id := c.Param("user"), c.Param("id")

	user, err := authorization.ExtractUser(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	if user != queried {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	var address model.Address
	if err := repository.DownloadAddress(id, &address); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNoContent, errors.New("no content"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, address)
}

func PostAddress(c *gin.Context) {
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	type Request struct {
		CEP        string `json:"cep"`
		Number     int    `json:"number"`
		Complement string `json:"complement"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
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
}
