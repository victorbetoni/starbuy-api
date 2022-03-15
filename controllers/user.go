package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"starbuy/authorization"
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

func QueryUser(w http.ResponseWriter, r *http.Request) {
	queried := mux.Vars(r)["username"]
	var user model.User

	if err := repository.DownloadUser(queried, &user); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNotFound, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

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
