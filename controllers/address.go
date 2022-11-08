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

func GetAddresses(c *gin.Context) (int, error) {
	user, _ := authorization.ExtractUser(c)

	var addresses []model.RawAddress
	if err := repository.DownloadAddresses(user, &addresses); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	for _, add := range addresses {
		if add.Holder != user {
			return http.StatusNoContent, errors.New("unauthorized")
		}
	}

	c.JSON(http.StatusOK, addresses)
	return http.StatusOK, nil
}

func GetAddress(c *gin.Context) (int, error) {
	id := c.Param("id")

	user, _ := authorization.ExtractUser(c)

	var address model.Address
	if err := repository.DownloadAddress(id, &address); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	if address.Holder.Username != user {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	c.JSON(http.StatusOK, address)
	return http.StatusOK, nil
}

func PostAddress(c *gin.Context) (int, error) {
	user, _ := authorization.ExtractUser(c)

	type Request struct {
		CEP        string `json:"cep"`
		Number     int    `json:"number"`
		Complement string `json:"complement"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	req.CEP = strings.Replace(req.CEP, "-", "", 1)

	if len(req.CEP) > 8 {
		return http.StatusBadRequest, errors.New("bad request")
	}

	//TODO: Usar alguma API para verificar se o CEP bate com algum existente

	address := model.RawAddress{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		Holder:     user,
		CEP:        req.CEP,
		Number:     req.Number,
		Complement: req.Complement,
	}

	if err := repository.InsertAddress(address); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Endereço criado"})
	return http.StatusOK, nil
}
