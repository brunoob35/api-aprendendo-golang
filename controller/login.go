package controller

import (
	"Api-Aula1/auth"
	"Api-Aula1/models"
	"Api-Aula1/persistency"
	"Api-Aula1/repository"
	"Api-Aula1/responses"
	"Api-Aula1/security"
	"encoding/json"
	"io"
	"net/http"
)

// Login autentica um User na aplicação verificando o e-mail e senha do User
func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.Users
	if err := json.Unmarshal(reqBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repository.NewUsersRepo(db)
	userSalvoemDB, err := repo.FetchByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ValidatePassword(userSalvoemDB.Password, user.Password); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}
	
	token, err := auth.GenerateToken(uint64(userSalvoemDB.ID))
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}

	w.Write([]byte(token))
}
