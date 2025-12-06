package middlewares

import (
	"Api-Aula1/auth"
	"Api-Aula1/responses"
	"log"
	"net/http"
)

// Logger escreve informações da req no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}

// Autenticate verifica se o User está autenticado
func Autenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Err(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}
