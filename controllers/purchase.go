package controllers

import (
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
)

func QueryPurchases(w http.ResponseWriter, r *http.Request) {
	user, err := authorization.ExtractUser(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	var purchases []model.Purchase
	repository.DownloadPurchases(user, &purchases)

	responses.JSON(w, http.StatusOK, purchases)
}

func PostPurchase(w http.ResponseWriter, r *http.Request) {
	var err error
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.JSON(w, http.StatusUnprocessableEntity, err)
		return
	}

	var purchase model.RawPurchase
	if err = json.Unmarshal(body, &purchase); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	purchase.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)

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
	err = repository.DownloadItem(purchase.Item, &item)
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

	final := model.Purchase{
		Identifier: purchase.Identifier,
		Seller:     seller,
		Customer:   customer,
		Item:       item,
		Price:      item.Item.Price * float64(purchase.Quantity),
		Quantity:   purchase.Quantity,
	}

	repository.InsertPurchase(final)
	responses.JSON(w, http.StatusOK, final)
}
