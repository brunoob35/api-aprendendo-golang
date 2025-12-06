package models

import (
	"Api-Aula1/security"
	"Api-Aula1/utils"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type Users struct {
	ID       int8   `json:"id"`
	Name     string `json:"nome_usuario"`
	CPF      string `json:"cpf"`
	Email    string `json:"email_usuario"`
	Password string `json:"senha"`
}

func (u *Users) Prepare(step string) error {
	// Chama o validate()
	// Chama o format()
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *Users) validate(step string) error {
	//Aqui vamos verificar se os campos recebidos do usuário, não estão vazios!
	// Validar se o cpf é valido
	if u.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if u.Email == "" {
		return errors.New("O e-mail é obrigatório e não pode estar em branco")
	}

	// go get github.com/badoux/checkmail
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if err := utils.CPFValidator(u.CPF); err != nil {
		return err
	}

	if step == "create" && u.Password == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

func (u *Users) format(step string) error {
	// Aqui vamos formatar as Strings, para remover espaços
	// Depois também vamos aplicar hash na senha
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	u.Name = strings.ToLower(u.Name)
	u.Email = strings.ToLower(u.Email)

	// Todo: Chamada da criação do Hash de senha apenas quando estamos criando um novo User
	if step == "register" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		u.Password = string(hashedPassword)
	}

	return nil
}
