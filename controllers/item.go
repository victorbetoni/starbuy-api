package controllers

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"net/http"
	"os"
	"starbuy/authorization"
	"starbuy/model"
	"starbuy/repository"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostItem(c *gin.Context) (int, error) {

	var item model.PostedItem
	if err := c.BindJSON(&item); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	item.Item.Identifier = strings.Replace(uuid.New().String(), "-", "", 4)
	user, err := authorization.ExtractUser(c)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	resp, err := cld.Upload.Upload(c, item.Assets[0], uploader.UploadParams{
		PublicID: "assets/" + item.Item.Identifier})

	item.Assets[0] = resp.URL

	item.Item.Seller = user
	if err := repository.InsertItem(item); err != nil {
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, item)
	return 0, nil
}

func QueryItems(c *gin.Context) (int, error) {
	query := c.Param("query")

	var items []model.ItemWithAssets
	if err := repository.QueryItemsByName(query, &items); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}
	c.JSON(http.StatusOK, items)

	return 0, nil
}

func GetItem(c *gin.Context) (int, error) {
	queried := c.Param("id")
	key, ok := c.GetQuery("reviews")
	includeReviews := ok && key == "true"

	type ItemReview struct {
		User    model.User `json:"user,omitempty"`
		Message string     `json:"message"`
		Rate    int        `json:"rate"`
	}

	average := float64(-1)

	var incoming []model.Review
	var reviews []ItemReview
	if includeReviews {
		var err error
		average, err = repository.QueryProductReviews(queried, &incoming)
		if err != nil && err != sql.ErrNoRows {
			return http.StatusInternalServerError, err
		}
		for _, review := range incoming {
			reviews = append(reviews, ItemReview{User: review.User, Message: review.Message, Rate: review.Rate})
		}
	}

	var item model.ItemWithAssets
	if err := repository.DownloadItem(queried, &item); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("not found")
		}
		return http.StatusInternalServerError, err
	}

	type Response struct {
		Item    model.ItemWithAssets `json:"item,omitempty"`
		Reviews []ItemReview         `json:"reviews,omitempty"`
		Average float64              `json:"average"`
	}

	var response Response
	if includeReviews {
		response.Reviews = reviews
	}
	response.Item = item
	response.Average = average

	fmt.Println(response)

	c.JSON(http.StatusOK, response)
	return 0, nil
}

func GetAllItems(c *gin.Context) (int, error) {
	var items []model.ItemWithAssets
	if err := repository.DownloadAllItems(&items); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, items)
	return 0, nil
}

func GetCategory(c *gin.Context) (int, error) {
	queried, _ := strconv.Atoi(c.Param("id"))
	var items []model.ItemWithAssets

	if err := repository.DownloadItemByCategory(queried, &items); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNoContent, errors.New("no content")
		}
		return http.StatusInternalServerError, err
	}

	c.JSON(http.StatusOK, items)
	return 0, nil
}
