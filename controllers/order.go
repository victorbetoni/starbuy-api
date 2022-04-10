package controllers

import (
	"database/sql"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetPurchases(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)

	var purchases []model.Order
	repository.DownloadPurchases(user, &purchases)

	c.JSON(http.StatusOK, purchases)

	return nil
}

func GetPurchase(c *gin.Context) error {
	queried := c.Param("id")
	user, _ := authorization.ExtractUser(c)

	var purchase model.Order
	if err := repository.DownloadPurchase(queried, &purchase); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	if purchase.Customer.Username != user {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "unauthorized"})
		return nil
	}

	c.JSON(http.StatusOK, purchase)
	return nil
}

func PostPurchase(c *gin.Context) error {

	user, _ := authorization.ExtractUser(c)

	type Request struct {
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}

	req := Request{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	var customer model.User
	if err := repository.DownloadUser(user, &customer); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	var item model.ItemWithAssets
	if err := repository.DownloadItem(req.Item, &item); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	var seller model.User
	if err := repository.DownloadUser(item.Item.Seller.Username, &seller); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
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

	return nil
}
