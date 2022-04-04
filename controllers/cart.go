package controllers

import (
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"

	"github.com/gin-gonic/gin"
)

func QueryCart(c *gin.Context) {
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	var items []model.CartItem
	repository.DownloadCart(user, &items)
	c.JSON(http.StatusOK, items)
}

func PostCart(c *gin.Context) {

	type Request struct {
		Item     string `json:"item"`
		Quantity int    `json"quantity"`
	}

	req := Request{}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	user, err := authorization.ExtractUser(c)

	cart := model.RawCartItem{Holder: user, Quantity: req.Quantity, Item: req.Item}

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	cart.Holder = user
	repository.InsertCartItem(cart)
	c.JSON(http.StatusOK, cart)
}
