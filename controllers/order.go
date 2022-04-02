package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetPurchases(w http.ResponseWriter, r *http.Request) {
	user, err := authorization.ExtractUser(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	var purchases []model.Order
	repository.DownloadPurchases(user, &purchases)

	responses.JSON(w, http.StatusOK, purchases)
}

func GetPurchase(w http.ResponseWriter, r *http.Request) {
	queried := mux.Vars(r)["id"]
	user, err := authorization.ExtractUser(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	var purchase model.Order
	if err := repository.DownloadPurchase(queried, &purchase); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNoContent, errors.New("Compra não encontrada"))
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if purchase.Customer.Username != user {
		responses.Error(w, http.StatusUnauthorized, errors.New("Não autorizado"))
		return
	}

	responses.JSON(w, http.StatusOK, purchase)
}

func PostPurchase(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		Item     string `json:"item"`
		Quantity int    `json:"quantity"`
	}

	var err error
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var req Request
	if err = json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err := authorization.ExtractUser(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	var customer model.User
	err = repository.DownloadUser(user, &customer)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var item model.ItemWithAssets
	err = repository.DownloadItem(req.Item, &item)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var seller model.User
	err = repository.DownloadUser(item.Item.Seller.Username, &seller)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	final := model.Order{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		Seller:     seller,
		Customer:   customer,
		Item:       item,
		Price:      item.Item.Price * float64(req.Quantity),
		Quantity:   req.Quantity,
	}

	repository.InsertPurchase(final)
	responses.JSON(w, http.StatusOK, final)
}
