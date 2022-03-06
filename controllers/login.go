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
		return
	}

	var login database.Login
	if err = json.Unmarshal(body, &login); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db := database.GrabDB()
	var recorded database.Login
	if err = db.Get(&recorded, "SELECT * FROM login WHERE username=$1", login.Username); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ComparePassword(recorded.Password, login.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	responses.JSON(w, http.StatusOK, err)

}

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var data database.IncomingUser
	if err = json.Unmarshal(body, &data); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db := database.GrabDB()

	testQueries := map[string]string{
		fmt.Sprintf("SELECT * FROM users WHERE username='%s'", data.Username): "Username já está em uso",
		fmt.Sprintf("SELECT * FROM users WHERE email='%s'", data.Email):       "Email já está em uso",
	}

	var found database.User
	for key, value := range testQueries {
		err := db.Get(&found, key)
		if (err != nil && err != sql.ErrNoRows) || err == nil {
			responses.Error(w, http.StatusBadRequest, errors.New(value))
			return
		}
	}

	tx := db.MustBegin()
	user := database.User{data.Username, data.Email, data.Name, data.Gender, data.Registration, data.Birthdate, data.Seller}
	tx.NamedExec("INSERT INTO users VALUES (:username,:email,:name,:gender,:registration,:birthdate,:seller)", &user)
	if err := tx.Commit(); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	crypt, err := security.Hash(data.Password)
	tx2 := db.MustBegin()
	tx2.MustExec("INSERT INTO login VALUES ($1,$2)", data.Username, string(crypt))
	if err := tx2.Commit(); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, nil)
}
