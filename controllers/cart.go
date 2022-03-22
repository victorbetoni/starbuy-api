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
)

func QueryCart(w http.ResponseWriter, r *http.Request) {
	user, err := authorization.ExtractUser(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	var items []model.CartItem
	repository.DownloadCart(user, &items)
	responses.JSON(w, http.StatusOK, items)
}

func PostCart(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var cart model.RawCartItem
	if err = json.Unmarshal(body, &cart); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	user, err := authorization.ExtractUser(r)

	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token invalido"))
		return
	}

	cart.Holder = user
	repository.InsertCartItem(cart)
	responses.JSON(w, http.StatusOK, cart)
}
