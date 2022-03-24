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
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func GetReviews(w http.ResponseWriter, r *http.Request) {
	queried := mux.Vars(r)["user"]

	var reviews []model.Review
	if err := repository.QueryUserReviews(queried, &reviews); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNoContent, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, reviews)
}

func GetReview(w http.ResponseWriter, r *http.Request) {
	queried := mux.Vars(r)["id"]

	var review model.Review
	if err := repository.DownloadReview(queried, &review); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNoContent, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, review)
}

func PostReview(w http.ResponseWriter, r *http.Request) {

	type Request struct {
		Message string `json:"message"`
		Rate    int    `json:"rate"`
		Item    string `json:"item"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var req Request
	if err = json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if req.Rate > 10 {
		req.Rate = 10
	}

	if req.Rate < 0 {
		req.Rate = 0
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

	final := model.RawReview{
		Identifier: strings.Replace(uuid.New().String(), "-", "", 4),
		User:       username,
		Item:       req.Item,
		Message:    req.Message,
		Rate:       req.Rate,
	}

	repository.InsertReview(final)

	responses.JSON(w, http.StatusOK, final)
}

func PutReview(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Review  string `json:"id"`
		Rate    int    `json:"rate"`
		Message string `json:"message"`
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var req Request
	if err = json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	username, erro := authorization.ExtractUser(r)
	if erro != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	var review model.Review
	if err := repository.DownloadReview(req.Review, &review); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if review.User.Username != username {
		responses.Error(w, http.StatusUnauthorized, errors.New("Não autorizado"))
		return
	}

	final := model.RawReview{Identifier: review.Identifier, User: username, Message: req.Message, Rate: req.Rate, Item: review.Item.Item.Identifier}

	repository.UpdateReview(final)
	responses.JSON(w, http.StatusOK, final)
}

func DeleteReview(w http.ResponseWriter, r *http.Request) {

}
