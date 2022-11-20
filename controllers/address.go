package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
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
		Name       string `json:"name"`
		CEP        string `json:"cep"`
		Number     int    `json:"number"`
		Complement string `json:"complement"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	type Count struct {
		Count int `db:"count"`
	}

	var count Count
	db := database.GrabDB()
	if err := db.Get(&count, "SELECT COUNT(*) FROM address WHERE name=$1 AND holder=$2", req.Name, user); err != nil {
		return http.StatusInternalServerError, err
	}

	if count.Count != 0 {
		return http.StatusBadRequest, errors.New("Você já tem um endereço com esse nome")
	}

	req.CEP = strings.Replace(req.CEP, "-", "", 1)

	if len(req.CEP) > 8 {
		return http.StatusBadRequest, errors.New("bad request")
	}

	address := model.RawAddress{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		Holder:     user,
		CEP:        req.CEP,
		Number:     req.Number,
		Complement: req.Complement,
		Name:       req.Name,
	}
	if err := repository.InsertAddress(address); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Endereço criado"})
	return http.StatusOK, nil
}

func DeleteAddress(c *gin.Context) (int, error) {
	id := c.Param("id")

	user, _ := authorization.ExtractUser(c)

	if err := repository.DeleteAddress(id, user); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Endereço removido"})
	return 0, nil
}
