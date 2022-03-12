package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"
	"starbuy/security"
)

type Login struct {
	Username string `db:"username"`
	Password string `db:"password"`
}

type IncomingUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Gender         int    `json:"gender"`
	Registration   string `json:"registration"`
	Birthdate      string `json:"birthdate"`
	Seller         bool   `json:"seller"`
	Password       string `json:"password"`
	ProfilePicture string `json:"profile_picture"`
	City           string `json:"city"`
}

func Auth(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var login Login
	if err = json.Unmarshal(body, &login); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db := database.GrabDB()
	var recorded Login
	if err = db.Get(&recorded, "SELECT * FROM login WHERE username=$1", login.Username); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	var user model.User
	if err = repository.DownloadUser(login.Username, &user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ComparePassword(recorded.Password, login.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token := authorization.GenerateToken(login.Username)

	type Response struct {
		User  model.User `json:"user"`
		Token string     `json:"jwt"`
	}

	responses.JSON(w, http.StatusOK, Response{User: user, Token: token})

}
