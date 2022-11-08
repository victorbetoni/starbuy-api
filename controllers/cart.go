package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"

	"github.com/gin-gonic/gin"
)

func QueryCart(c *gin.Context) (int, error) {
	user, _ := authorization.ExtractUser(c)

	var items []model.CartItem
	if err := repository.DownloadCart(user, &items); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}
	c.JSON(http.StatusOK, items)
	return 0, nil
}

func DeleteCart(c *gin.Context) (int, error) {
	item := c.Param("item")
	username, _ := authorization.ExtractUser(c)

	if err := repository.DeleteFromCart(username, item); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Item removido do carrinho"})
	return 0, nil
}

func PostCart(c *gin.Context) (int, error) {

	type Request struct {
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	username, _ := authorization.ExtractUser(c)
	cart := model.RawCartItem{Holder: username, Quantity: req.Quantity, Item: req.Item}

	var item model.ItemWithAssets
	if err := repository.DownloadItem(req.Item, &item); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("item not found")
		}
		return http.StatusInternalServerError, err
	}

	if item.Item.Stock < cart.Quantity {
		return http.StatusUnauthorized, errors.New("Estoque insuficiente")
	}

	var user model.User
	if err := repository.DownloadUser(username, &user); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("user not found")
		}
		return http.StatusInternalServerError, err
	}

	cart.Holder = username
	if err := repository.InsertCartItem(cart); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Item adicionado ao carrinho!"})
	return 0, nil
}
