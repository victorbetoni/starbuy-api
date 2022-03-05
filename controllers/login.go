package controllers

import (
	"authentication-service/database"
	"authentication-service/responses"
	"authentication-service/security"
	"encoding/json"
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
	if err = db.Get(&recorded, "SELECT * FROM login WHERE username=?", login.Username); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	}

	if err = security.ComparePassword(recorded.Password, login.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
	}

	responses.JSON(w, http.StatusOK, err)

}
