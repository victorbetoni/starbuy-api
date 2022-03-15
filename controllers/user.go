package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"
	"starbuy/security"

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

	db := database.GrabDB()

	testQueries := map[string]string{
		fmt.Sprintf("SELECT * FROM users WHERE username='%s'", data.Username): "Username já está em uso",
		fmt.Sprintf("SELECT * FROM users WHERE email='%s'", data.Email):       "Email já está em uso",
	}

	var found model.User
	for key, value := range testQueries {
		err := db.Get(&found, key)
		if (err != nil && err != sql.ErrNoRows) || err == nil {
			responses.Error(w, http.StatusBadRequest, errors.New(value))
			return
		}
	}

	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO users VALUES (:username,:email,:name,:gender,:registration,:birthdate,:seller,:profile_picture,:city)", &user)
	if err := tx.Commit(); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	crypt, _ := security.Hash(data.Password)
	tx2 := db.MustBegin()
	tx2.MustExec("INSERT INTO login VALUES ($1,$2)", data.Username, string(crypt))
	if err := tx2.Commit(); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
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
