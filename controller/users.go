package controller

import (
	"log"
	"net/http"
)

func CreateUser(writer http.ResponseWriter, request *http.Request) {
	// Ler body request
	// Fazer o Unmarshal do json na struct do User
	// Fazer os tratamentos necessários usando os métodos de User
	// Enviar para o Repositório (por hora pode pular este passo, vamos trabalhar a persistencia na próxima aula)
	// Preparar a resposta (por hora pode retornar apenas os dados do seu user ou um ID aleatório)
	log.Print("CreateUser")
}

func FetchUser(writer http.ResponseWriter, request *http.Request) {

}

func UpdateUser(writer http.ResponseWriter, request *http.Request) {

}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {

}
