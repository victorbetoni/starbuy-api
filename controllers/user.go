package controllers

import (
	"database/sql"
	"errors"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"net/http"
	"os"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type IncomingUser struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Birthdate      string `json:"birthdate"`
	Seller         bool   `json:"seller"`
	ProfilePicture string `json:"profile_picture"`
	City           string `json:"city"`
	Password       string `json:"password"`
}

func Register(c *gin.Context) (int, error) {

	incoming := IncomingUser{}
	if err := c.BindJSON(&incoming); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	user := model.User{
		Username:       incoming.Username,
		Email:          incoming.Email,
		Name:           incoming.Name,
		Birthdate:      incoming.Birthdate,
		ProfilePicture: incoming.ProfilePicture,
		Seller:         incoming.Seller,
		City:           incoming.City,
		Registration:   time.Now().Format("2006-01-02"),
	}

	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	resp, err := cld.Upload.Upload(c, user.ProfilePicture, uploader.UploadParams{
		PublicID: "profile_pic/" + user.Username})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.ProfilePicture = resp.URL

	if err := repository.InsertUser(user, incoming.Password); err != nil {
		return http.StatusInternalServerError, err
	}

	token := authorization.GenerateToken(user.Username)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Registrado com sucesso", "user": user, "jwt": token})
	return 0, nil
}

func PostUserProfilePicture(c *gin.Context) (int, error) {
	type Body struct {
		Image string `json:"imageB64"`
	}

	incoming := Body{}
	if err := c.BindJSON(&incoming); err != nil {
		return http.StatusBadRequest, errors.New("bad request")
	}

	username, err := authorization.ExtractUser(c)
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	resp, err := cld.Upload.Upload(c, incoming.Image, uploader.UploadParams{
		PublicID: "profile_pic/" + username})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	db := database.GrabDB()
	tx := db.MustBegin()
	tx.MustExec("UPDATE users SET profile_picture=$1 WHERE username=$2", resp.URL, username)
	if err := tx.Commit(); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func GetUser(c *gin.Context) (int, error) {
	queried := c.Param("user")

	var user model.User

	type Response struct {
		User    model.User             `json:"user,omitempty"`
		Items   []model.ItemWithAssets `json:"items,omitempty"`
		Reviews []model.Review         `json:"reviews,omitempty"`
		Rating  float64                `json:"rating"`
	}

	response := Response{}
	if err := repository.DownloadUser(queried, &user); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Usuário não encontrado")
		}
		return http.StatusInternalServerError, err
	}
	response.User = user

	if key, ok := c.GetQuery("includeItems"); ok && key == "true" {
		var local []model.ItemWithAssets
		if err := repository.DownloadUserProducts(queried, &local); err != nil && err != sql.ErrNoRows {
			return http.StatusInternalServerError, err
		}

		var items []model.ItemWithAssets
		//Removing seller (duplicated data)
		for _, item := range local {
			final := model.Item{
				Identifier:  item.Item.Identifier,
				Title:       item.Item.Title,
				Category:    item.Item.Category,
				Stock:       item.Item.Stock,
				Description: item.Item.Description,
				Price:       item.Item.Price,
			}
			items = append(items, model.ItemWithAssets{Item: final, Assets: item.Assets})
		}
		response.Items = items
	}

	if key, ok := c.GetQuery("includeReviews"); ok && key == "true" {
		var reviews []model.Review
		var local []model.Review
		var average float64
		if loc, err := repository.QueryUserReceivedReviews(queried, &local); err != nil && err != sql.ErrNoRows {
			average = loc
			return http.StatusInternalServerError, err
		}

		//Removing reviewer (duplicated data)
		for _, review := range local {
			final := model.Review{
				Message: review.Message,
				Item:    review.Item,
				Rate:    review.Rate,
			}
			reviews = append(reviews, final)
		}
		response.Rating = average
		response.Reviews = local
	}

	if err := repository.DownloadUser(queried, &user); err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound, errors.New("Usuário não encontrado")
		}
		return http.StatusInternalServerError, err
	}

	response.User = user
	c.JSON(http.StatusOK, response)

	return http.StatusOK, nil
}
