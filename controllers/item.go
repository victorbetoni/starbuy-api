package controllers

import (
	"database/sql"
	"errors"
	"net/http"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostItem(c *gin.Context) {

	var item model.PostedItem
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	item.Item.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	item.Item.Seller = user
	repository.InsertItem(item)
	c.JSON(http.StatusOK, item)

}

func GetItem(c *gin.Context) {
	queried := c.Param("id")
	key, ok := c.GetQuery("reviews")
	includeReviews := ok && key == "true"

	var reviews []model.Review
	if includeReviews {
		if err := repository.QueryProductReviews(queried, &reviews); err != nil && err != sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, review := range reviews {
			review.Item = model.ItemWithAssets{}
		}
	}

	var item model.ItemWithAssets
	if err := repository.DownloadItem(queried, &item); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNoContent, err)
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	type Response struct {
		Item    *model.ItemWithAssets `json:"item,omitempty"`
		Reviews *[]model.Review       `json:"reviews,omitempty"`
	}

	var response Response
	if includeReviews {
		response.Reviews = &reviews
	}
	response.Item = &item

	c.JSON(http.StatusOK, response)
}

func GetAllItems(c *gin.Context) {
	var items []model.ItemWithAssets
	if err := repository.DownloadAllItems(&items); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNoContent, errors.New("no content"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, items)
}

func GetCategory(c *gin.Context) {
	queried, _ := strconv.Atoi(c.Param("id"))
	var items []model.ItemWithAssets

	if err := repository.DownloadItemByCategory(queried, &items); err != nil {
		if err == sql.ErrNoRows {
			c.AbortWithError(http.StatusNoContent, errors.New("no content"))
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, items)
}
