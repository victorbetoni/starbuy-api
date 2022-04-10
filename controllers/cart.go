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

func QueryCart(c *gin.Context) error {
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return nil
	}

	var items []model.CartItem
	repository.DownloadCart(user, &items)
	c.JSON(http.StatusOK, items)
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "quantity greater than item stock"})
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

	c.JSON(http.StatusOK, model.CartItem{Holder: user, Quantity: cart.Quantity, Item: item.Item})
	return nil
}
