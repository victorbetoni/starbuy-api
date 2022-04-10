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

func PostItem(c *gin.Context) error {

	var item model.PostedItem
	if err := c.BindJSON(&item); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	item.Item.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)
	user, err := authorization.ExtractUser(c)

	if err != nil {
		c.Error(err)
		c.AbortWithError(http.StatusUnauthorized, errors.New("invalid token"))
		return nil
	}

	item.Item.Seller = user
	repository.InsertItem(item)
	c.JSON(http.StatusOK, item)
	return nil
}

func GetItem(c *gin.Context) error {
	queried := c.Param("id")
	key, ok := c.GetQuery("reviews")
	includeReviews := ok && key == "true"

	var reviews []model.Review
	if includeReviews {
		if err := repository.QueryProductReviews(queried, &reviews); err != nil && err != sql.ErrNoRows {
			return err
		}
		for _, review := range reviews {
			review.Item = model.ItemWithAssets{}
		}
	}

	var item model.ItemWithAssets
	if err := repository.DownloadItem(queried, &item); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
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
	return nil
}

func GetAllItems(c *gin.Context) error {
	var items []model.ItemWithAssets
	if err := repository.DownloadAllItems(&items); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}
	c.JSON(http.StatusOK, items)

	return nil
}

func GetCategory(c *gin.Context) error {
	queried, _ := strconv.Atoi(c.Param("id"))
	var items []model.ItemWithAssets

	if err := repository.DownloadItemByCategory(queried, &items); err != nil {
		if err == sql.ErrNoRows {
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": false, "message": "no content"})
			return nil
		}
		return err
	}

	c.JSON(http.StatusOK, items)
	return nil
}
