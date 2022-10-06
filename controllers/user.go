package controllers

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"starbuy/authorization"
	"starbuy/database"
	"starbuy/model"
	"starbuy/repository"
	"strings"
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

func Register(c *gin.Context) error {

	incoming := IncomingUser{}
	if err := c.BindJSON(&incoming); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
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

	fmt.Println(user)

	if err := user.Prepare(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error(), "user": nil, "jwt": ""})
		return nil
	}

	if err := repository.InsertUser(user, incoming.Password); err != nil {
		return nil
	}

	token := authorization.GenerateToken(user.Username)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Registrado com sucesso", "user": user, "jwt": token})
	return nil
}

func writeImage(fileNameBase, data string) (io.Reader, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return nil, errors.New("Imagem inválida")
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return nil, err
	}
	_, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return nil, err
	}

	fileName := fileNameBase + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return reader, nil
}

func PostUserProfilePicture(c *gin.Context) error {
	type Body struct {
		Image string `json:"imageB64"`
	}

	incoming := Body{}
	if err := c.BindJSON(&incoming); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": false, "message": "bad request"})
		return nil
	}

	file, err := writeImage("img", incoming.Image)
	if err != nil {
		return err
	}

	username, err := authorization.ExtractUser(c)
	cld, _ := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))
	resp, err := cld.Upload.Upload(c, file, uploader.UploadParams{
		PublicID: "profile_pic/" + username})

	fmt.Println(resp.SecureURL)
	fmt.Println(resp.Error.Message)

	if err != nil {
		return err
	}

	db := database.GrabDB()
	tx := db.MustBegin()
	tx.MustExec("UPDATE users SET profile_picture=$1 WHERE username=$2", resp.URL, username)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func GetUser(c *gin.Context) error {
	queried := c.Param("user")

	includeItems, includeReviews := false, false

	if key, ok := c.GetQuery("includeItems"); ok && key == "true" {
		includeItems = true
	}
	if key, ok := c.GetQuery("includeReviews"); ok && key == "true" {
		includeReviews = true
	}

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
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}
	response.User = user

	if includeItems {
		var local []model.ItemWithAssets
		if err := repository.DownloadUserProducts(queried, &local); err != nil && err != sql.ErrNoRows {
			return err
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

	if includeReviews {
		var reviews []model.Review
		var local []model.Review
		var average float64
		if loc, err := repository.QueryUserReceivedReviews(queried, &local); err != nil && err != sql.ErrNoRows {
			average = loc
			return err
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
			c.Error(err)
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": false, "message": "not found"})
			return nil
		}
		return err
	}

	response.User = user

	fmt.Println(user)
	c.JSON(http.StatusOK, response)
	return nil
}
