package controllers

import (
	"authentication-service/model"
	"authentication-service/repository"
	"authentication-service/responses"
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
)

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
