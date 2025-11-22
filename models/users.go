package models

type Users struct {
	id       int8
	name     string
	email    string
	cpf      string
	password string
}

func (u *Users) Prepare() error {
	// Chama o validate()
	// Chama o format()
	return nil
}

func (u *Users) validate(step string) error {
	//Aqui vamos verificar se os campos recebidos do usuário, não estão vazios!
	// Validar se o cpf é valido
	return nil
}

func (u *Users) format(step string) error {
	// Aqui vamos formatar as Strings, para remover espaços
	// Depois também vamos aplicar hash na senha
	return nil
}
