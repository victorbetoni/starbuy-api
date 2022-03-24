package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"
)

func PostReview(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var review model.RawReview
	if err = json.Unmarshal(body, &review); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	username, erro := authorization.ExtractUser(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	db := database.GrabDB()
	if err := db.Get(nil, "SELECT * FROM purchase_log WHERE holder=$1 AND product=$2", username, review.Item); err != nil && err == sql.ErrNoRows {
		responses.Error(w, http.StatusUnauthorized, errors.New("Você não comprou esse produto"))
		return
	}

	review.User = username
	repository.InsertReview(review)

	responses.JSON(w, http.StatusOK, review)
}
