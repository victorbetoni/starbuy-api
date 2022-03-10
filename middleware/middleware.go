package middleware

import (
	"fmt"
	"net/http"
)

func Authorize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verificando credenciais")
		next(w, r)
	}
}
