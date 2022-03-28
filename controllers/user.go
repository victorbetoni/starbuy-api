package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"

	"github.com/gorilla/mux"
)

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var data IncomingUser
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	user := model.User{
		Username:       data.Username,
		Email:          data.Email,
		Name:           data.Name,
		Gender:         data.Gender,
		Birthdate:      data.Birthdate,
		Seller:         data.Seller,
		ProfilePicture: data.ProfilePicture,
		City:           data.City,
		Registration:   data.Registration}

	if err := user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err := repository.InsertUser(user, data.Password); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	queried := mux.Vars(r)["username"]
	var user model.User

	keys, ok := r.URL.Query()["includeItems"]

	if !ok || len(keys[0]) < 1 {
		responses.Error(w, http.StatusNotFound, errors.New("Missing key"))
		return
	}

	key := keys[0]

	var items []model.ItemWithAssets

	if key == "true" {
		var local []model.ItemWithAssets
		if err := repository.DownloadUserProducts(queried, &local); err != nil {
			if err == sql.ErrNoRows {
				responses.Error(w, http.StatusNotFound, err)
				return
			}
			responses.Error(w, http.StatusInternalServerError, err)
			return
		}

		//Removing seller (duplicated data)
		for _, item := range local {
			final := model.Item{
				Identifier:  item.Item.Identifier,
				Title:       item.Item.Title,
				Category:    item.Item.Category,
				Stock:       item.Item.Stock,
				Description: item.Item.Description,
				Price:       item.Item.Price,
			}
			items = append(items, model.ItemWithAssets{Item: final, Assets: item.Assets})
		}
	}

	if err := repository.DownloadUser(queried, &user); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNotFound, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if key == "true" {

		type UserWithItem struct {
			User  model.User             `json:"user"`
			Items []model.ItemWithAssets `json:"items"`
		}
		responses.JSON(w, http.StatusOK, UserWithItem{user, items})
		return
	}

	responses.JSON(w, http.StatusOK, user)
}
