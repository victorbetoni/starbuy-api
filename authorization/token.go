package authorization

import (
	"errors"
	"fmt"
	"log"
	"starbuy/util"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(username string) string {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var config = util.GrabConfig()
	fmt.Println(config.Secret)
	str, err := token.SignedString([]byte(config.Secret))

	if err != nil {
		log.Fatal(err)
	}

	return str
}

func ValidateToken(c *gin.Context) error {
	raw := extractToken(c)

	token, err := jwt.Parse(raw, checkSecurityKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func extractToken(c *gin.Context) string {
	raw := c.GetHeader("Authorization")

	if len(strings.Split(raw, " ")) != 2 {
		return ""
	}
	return strings.Split(raw, " ")[1]
}

func checkSecurityKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
	}

	return []byte(util.GrabConfig().Secret), nil
}

func ExtractUser(c *gin.Context) (string, error) {
	raw := extractToken(c)
	token, err := jwt.Parse(raw, checkSecurityKey)
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := fmt.Sprintf("%v", claims["username"])
		return username, nil
	}

	return "", errors.New("Invalid token")
}
