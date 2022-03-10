package controllers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"starbuy/model"
	"starbuy/repository"
	"starbuy/responses"
	"strconv"

	"github.com/gorilla/mux"
)

func PostItem(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var item model.ItemWithAssets
	if err = json.Unmarshal(body, &item); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	repository.InsertItem(item)

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
