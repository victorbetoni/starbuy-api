package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func PostItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var item model.PostedItem
	if err = json.Unmarshal(body, &item); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	item.Item.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)

	user, err := authorization.ExtractUser(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, errors.New("Token inválido"))
		return
	}

	item.Item.Seller = user
	repository.InsertItem(item)
	responses.JSON(w, http.StatusOK, item)

}

func QueryItem(w http.ResponseWriter, r *http.Request) {
	queried := mux.Vars(r)["id"]
	var item model.ItemWithAssets

	if err := repository.DownloadItem(queried, &item); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNotFound, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, item)
}

func QueryAllItems(w http.ResponseWriter, r *http.Request) {
	var items []model.ItemWithAssets
	if err := repository.DownloadAllItems(&items); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNotFound, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, items)
}

func QueryCategory(w http.ResponseWriter, r *http.Request) {
	queried, _ := strconv.Atoi(mux.Vars(r)["id"])
	var items []model.ItemWithAssets

	if err := repository.DownloadItemByCategory(queried, &items); err != nil {
		if err == sql.ErrNoRows {
			responses.Error(w, http.StatusNotFound, err)
			return
		}
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, items)
}
