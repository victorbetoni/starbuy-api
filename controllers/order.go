package controllers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetPurchases(c *gin.Context) {
	user, err := authorization.ExtractUser(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	var purchases []model.Order
	repository.DownloadPurchases(user, &purchases)

	c.JSON(http.StatusOK, purchases)
}

func GetPurchase(c *gin.Context) {
	queried := c.Param("id")
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	var purchase model.Order
	if err := repository.DownloadPurchase(queried, &purchase); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNotFound, errors.New("not found"))
			return
		}
		log.Fatal(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if purchase.Customer.Username != user {
		c.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	c.JSON(http.StatusOK, purchase)
}

func PostPurchase(c *gin.Context) {

	user, err := authorization.ExtractUser(c)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	type Request struct {
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}

	req := Request{}
	var customer model.User
	err = repository.DownloadUser(user, &customer)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var item model.ItemWithAssets
	err = repository.DownloadItem(req.Item, &item)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var seller model.User
	err = repository.DownloadUser(item.Item.Seller.Username, &seller)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := c.BindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	final := model.Order{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		Seller:     seller,
		Customer:   customer,
		Item:       item,
		Price:      float64(item.Item.Price * (float64)(req.Quantity)),
		Quantity:   req.Quantity,
	}

	repository.InsertPurchase(final)
	c.JSON(http.StatusOK, final)
}
