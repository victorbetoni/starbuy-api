package controllers

import (
	"authentication-service/model"
	"authentication-service/repository"
	"authentication-service/responses"
	"database/sql"
	"fmt"
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
		fmt.Println("EXPLODIU AQUI HEIN")
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, item)
}
