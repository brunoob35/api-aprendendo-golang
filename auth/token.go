package auth

import (
	"Api-Aula1/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go" // Import com alias só por causa do traço no nome
)

// go get "github.com/dgrijalva/jwt-go"

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() // tempo para expirar o token
	permissions["userId"] = userID

	// jwt.SigningMethodHS256 é um de muitos diferentes métodos para assinatura do token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString(config.SecretKey)
}

// ValidateToken recebe apenas o token para validação
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, retriveAuthKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid Token")
}

// extractToken is not going to be exported.
// Valida se há um token dentro do objeto bearer da requisição
// Então extrai apenas a parte do token e envia para validação.
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

// retriveAuthKey Retorna qual o metodo de autenticação dado um token
func retriveAuthKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Token signing method unexpected! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
