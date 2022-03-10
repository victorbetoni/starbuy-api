package authorization

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"starbuy/util"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

func ValidateToken(request *http.Request) error {
	raw := extractToken(request)

	token, err := jwt.Parse(raw, checkSecurityKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")
}

func extractToken(r *http.Request) string {
	raw := r.Header.Get("Authorization")

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

func ExtractUser(r *http.Request) (string, error) {
	raw := extractToken(r)
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
