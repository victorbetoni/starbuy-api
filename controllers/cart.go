package controllers

import (
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

	user, err := authorization.ExtractUser(c)

	cart := model.RawCartItem{Holder: user, Quantity: req.Quantity, Item: req.Item}

	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "invalid token"})
		return nil
	}

	cart.Holder = user
	repository.InsertCartItem(cart)
	c.JSON(http.StatusOK, cart)
	return nil
}
