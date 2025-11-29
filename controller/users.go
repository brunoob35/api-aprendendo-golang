package controller

import (
	"Api-Aula1/models"
	"Api-Aula1/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Ler body request
	// Fazer o Unmarshal do json na struct do User
	// Fazer os tratamentos necessários usando os métodos de User
	// Enviar para o Repositório (por hora pode pular este passo, vamos trabalhar a persistencia na próxima aula)
	// Preparar a resposta (por hora pode retornar apenas os dados do seu user ou um ID aleatório)

	//Lê o request.Body
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	// Descompacta (Unmarshal) o conteúdo JSON em uma Struct
	var newUser models.Users
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	// Chama os métodos de preparação do User
	if err = newUser.Prepare("create"); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	responses.JSON(w, http.StatusCreated, newUser)
}

func FetchUser(writer http.ResponseWriter, request *http.Request) {

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {}
