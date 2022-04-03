package controllers

import (
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

func QueryCart(c *gin.Context) {
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	var items []model.CartItem
	repository.DownloadCart(user, &items)
	c.JSON(http.StatusOK, items)
}

func PostCart(c *gin.Context) {

	item := c.PostForm("item")
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))

	user, err := authorization.ExtractUser(c)

	cart := model.RawCartItem{Holder: user, Quantity: quantity, Item: item}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error:": "Token invalido"})
		return
	}

	cart.Holder = user
	repository.InsertCartItem(cart)
	c.JSON(http.StatusOK, cart)
}
