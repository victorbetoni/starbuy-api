package controllers

import (
	"authentication-service/database"
	"authentication-service/responses"
	"authentication-service/security"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}

	var login database.Login
	if err = json.Unmarshal(body, &login); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db := database.GrabDB()
	var recorded database.Login
	if err = db.Get(&recorded, "SELECT * FROM login WHERE username=$1", login.Username); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	if err = security.ComparePassword(recorded.Password, login.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
	}

	responses.JSON(w, http.StatusOK, err)

}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
	}

	var data database.PasswordBindedUser
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	}

	db := database.GrabDB()

	testQueries := map[string]string{
		fmt.Sprintf("SELECT * FROM users WHERE username=%s", data.Username): "Username já está em uso",
		fmt.Sprintf("SELECT * FROM users WHERE email=%s", data.Email):       "Email já está em uso",
	}

	for key, value := range testQueries {
		err := db.Get(nil, key)
		if (err != nil && err != sql.ErrNoRows) || err == nil {
			responses.Error(w, http.StatusBadRequest, errors.New(value))
		}
	}

	tx := db.MustBegin()
	tx.MustExec("INSERT INTO users VALUES (:username, :email, :name, :gender, :registration, :gender, :birthdate, :seller)", &data)
	if err := tx.Commit(); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	responses.JSON(w, http.StatusCreated, nil)
}
