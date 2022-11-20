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
	"time"
)

func GetOrders(c *gin.Context) (int, error) {
	user, _ := authorization.ExtractUser(c)

	var purchases []model.Order
	if err := repository.DownloadPurchases(user, &purchases); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Nenhum pedido encontrado")
		}
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, purchases)

	return 0, nil
}

func GetOrder(c *gin.Context) (int, error) {
	queried := c.Param("id")
	user, _ := authorization.ExtractUser(c)

	var purchase model.Order
	if err := repository.DownloadPurchase(queried, &purchase); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("not found")
		}
		return http.StatusInternalServerError, err
	}

	if purchase.Customer.Username != user {
		return http.StatusUnauthorized, errors.New("unauthorized")
	}

	c.JSON(http.StatusOK, purchase)
	return 0, nil
}

func CreateOrder(c *gin.Context) (int, error) {

	user, _ := authorization.ExtractUser(c)

	var seller model.User
	var customer model.User

	type OrderReq struct {
		SendTo string `json:"send_to"`
	}

	req := &OrderReq{}
	if err := c.BindJSON(&req); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	if err := repository.DownloadUser(user, &customer); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Cliente não encontrado")
		}
		return http.StatusInternalServerError, err
	}

	address := model.Address{}

	if err := repository.DownloadAddress(req.SendTo, &address); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Endereço não encontrado")
		}
		return http.StatusInternalServerError, err
	}

	if err := repository.DownloadUser(user, &seller); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Vendedor não encontrado")
		}
		return http.StatusInternalServerError, err
	}

	var items []model.CartItem
	if err := repository.DownloadCart(user, &items); err != nil {
		return http.StatusInternalServerError, err
	}

	for _, cart := range items {
		order := model.Order{
			Identifier: uuid.New().String(),
			Quantity:   cart.Quantity,
			Item:       *cart.Item,
			Customer:   customer,
			Seller:     seller,
			Price:      cart.Item.Item.Price * float64(cart.Quantity),
			Date:       time.Now().Format("2006-01-02"),
			SendTo:     address,
		}
		if err := repository.InsertPurchase(order); err != nil {
			return http.StatusInternalServerError, err
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Compra realizada com sucesso!"})

	return 0, nil
}

func UpdateOrder(c *gin.Context) (int, error) {
	user, _ := authorization.ExtractUser(c)
	queried := c.Param("id")

	var order model.Order
	if err := repository.DownloadPurchase(queried, &order); err != nil {
		return http.StatusInternalServerError, err
	}

	if order.Seller.Username != user && order.Customer.Username != user {
		return http.StatusNotFound, errors.New("Vendedor não encontrado")
	}

	if order.Status == 2 {
		return http.StatusUnauthorized, errors.New("Venda já finalizada")
	}

	if order.Status == 1 && order.Seller.Username == user {
		return http.StatusUnauthorized, errors.New("Você não pode atualizar esse pedido")
	}
	if order.Status == 0 && order.Customer.Username == user {
		return http.StatusUnauthorized, errors.New("Você não pode atualizar esse pedido")
	}

	if err := repository.UpdateOrder(order); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func GetReceivedOrders(c *gin.Context) (int, error) {
	user, _ := authorization.ExtractUser(c)

	var orders []model.Order
	if err := repository.DownloadOrders(user, &orders); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("not found")
		}
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, orders)
	return 0, nil
}
