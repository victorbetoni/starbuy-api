package controllers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
)

func GetOrders(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)

	var purchases []model.Order
	err := repository.DownloadPurchases(user, &purchases)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, purchases)

	return nil
}

func GetOrder(c *gin.Context) error {
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

func CreateOrder(c *gin.Context) error {

	user, _ := authorization.ExtractUser(c)

	var seller model.User
	var customer model.User

	if err := repository.DownloadUser(user, &customer); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Cliente não encontrado"})
			return nil
		}
		return err
	}

	if err := repository.DownloadUser(user, &seller); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Vendedor não encontrado"})
			return nil
		}
		return err
	}

	var items []model.CartItem
	if err := repository.DownloadCart(user, &items); err != nil {
		return err
	}

	for _, cart := range items {
		order := model.Order{
			Identifier: uuid.New().String(),
			Quantity:   cart.Quantity,
			Item:       *cart.Item,
			Customer:   customer,
			Seller:     seller,
			Price:      cart.Item.Item.Price * float64(cart.Quantity),
		}
		if err := repository.InsertPurchase(order); err != nil {
			return err
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Compra realizada com sucesso!"})

	return nil
}

func UpdateOrder(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)
	queried := c.Param("id")

	var order model.Order
	if err := repository.DownloadPurchase(queried, &order); err != nil {
		return err
	}

	if order.Seller.Username != user && order.Customer.Username != user {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Vendedor não encontrado"})
		return errors.New("vendedor não encontrado")
	}

	if order.Status == 2 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "A venda ja foi finalizada."})
		return errors.New("Não autorizado")
	}

	if order.Status == 1 && order.Seller.Username == user {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Você não pode atualizar esse pedido."})
		return errors.New("Não autorizado")
	}
	if order.Status == 0 && order.Customer.Username == user {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "Você não pode atualizar esse pedido."})
		return errors.New("Não autorizado")
	}

	if err := repository.UpdateOrder(order); err != nil {
		return err
	}

	return nil
}

func GetReceivedOrders(c *gin.Context) error {
	user, _ := authorization.ExtractUser(c)

	var orders []model.OrderWithItem
	if err := repository.DownloadOrders(user, &orders); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, orders)
	return nil
}
