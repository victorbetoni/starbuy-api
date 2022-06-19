package controllers

import (
	"database/sql"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"

	"github.com/gin-gonic/gin"
)

func QueryCart(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)

	var items []model.CartItem
	repository.DownloadCart(user, &items)
	c.JSON(http.StatusOK, items)
	return nil
}

func DeleteCart(c *gin.Context) error {
	item := c.Param("item")
	username, _ := authorization.ExtractUser(c)

	if err := repository.DeleteFromCart(username, item); err != nil {
		return nil
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Item removido do carrinho"})
	return nil
}

func PostCart(c *gin.Context) error {

	type Request struct {
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	username, _ := authorization.ExtractUser(c)
	cart := model.RawCartItem{Holder: username, Quantity: req.Quantity, Item: req.Item}

	var item model.ItemWithAssets
	if err := repository.DownloadItem(req.Item, &item); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "item not found"})
			return nil
		}
		return err
	}

	if item.Item.Stock < cart.Quantity {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "Estoque insuficiente."})
		return nil
	}

	var user model.User
	if err := repository.DownloadUser(username, &user); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "user not found"})
			return nil
		}
		return err
	}

	cart.Holder = username
	repository.InsertCartItem(cart)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Item adicionado ao carrinho!"})
	return nil
}
