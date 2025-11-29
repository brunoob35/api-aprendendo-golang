package utils

import (
	"errors"
	"strconv"
	"strings"
)

func CPFValidator(cpf string) error {
	cpf = onlyDigits(cpf)

	if len(cpf) != 11 { // CPF obrigatoriamente tem 11 dígitos
		return errors.New("CPF deve ter 11 dígitos")
	}

	if !CheckAllEqual(cpf) { // se todos os dígitos são iguais, CPF é inválido
		return errors.New("números iguais, CPF inválido")
	}

	if !CalcularDv1(cpf) { // valida o primeiro dígito verificador
		return errors.New("DV1 inválido, CPF inválido")
	}

	if !CalcularDv2(cpf) { // valida o segundo dígito verificador
		return errors.New("DV2 inválido, CPF inválido")
	}

	return nil
}

// CalcularDv2 calcula o segundo dígito verificador (posição 10)
func CalcularDv2(cpf string) bool {
	digits := strings.Split(cpf, "") // separa em ["1","2",...]
	if len(digits) < 11 {
		return false
	}

	soma := 0
	// regra: multiplicar os 10 primeiros dígitos pelos pesos 11 a 2
	for i := 0; i < 10; i++ {
		n, err := strconv.Atoi(digits[i])
		if err != nil {
			return false
		}
		soma += n * (11 - i)
	}

	// fórmula igual ao DV1
	dv := (soma * 10) % 11
	if dv == 10 {
		dv = 0
	}

	return strconv.Itoa(dv) == digits[10] // compara com o dígito final
}

// CalcularDv1 calcula o primeiro dígito verificador (posição 9)
func CalcularDv1(cpf string) bool {
	digits := strings.Split(cpf, "") // separa em slice ["1","2","3",...]
	if len(digits) < 10 {            // precisa ter pelo menos 10 dígitos
		return false
	}

	soma := 0
	// regra: multiplicar os 9 primeiros dígitos pelos pesos 10 a 2
	for i := 0; i < 9; i++ {
		n, err := strconv.Atoi(digits[i]) // converte string pra inteiro
		if err != nil {
			return false
		}
		soma += n * (10 - i) // aplica peso
	}

	// fórmula do DV: (soma * 10) % 11
	dv := (soma * 10) % 11
	if dv == 10 { // DV vira 0 se der 10
		dv = 0
	}

	return strconv.Itoa(dv) == digits[9] // compara com o dígito do CPF
}

// CheckAllEqual retorna true se o CPF NÃO for composto só de um número repetido
// Ex: "11111111111", "00000000000" → inválido
func CheckAllEqual(cpf string) bool {
	if len(cpf) == 0 {
		return false
	}

	first := cpf[0] // pega o primeiro caractere
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != first { // encontrou número diferente
			return true // então NÃO são todos iguais → válido
		}
	}
	return false // todos iguais → inválido
}

// onlyDigits remove qualquer coisa que não seja número
// Ex: "123.456.789-10" vira "12345678910"
func onlyDigits(s string) string {
	var b strings.Builder
	b.Grow(len(s)) // otimiza alocação

	for _, r := range s {
		if r >= '0' && r <= '9' { // mantém só números
			b.WriteRune(r)
		}
	}
	return b.String()
}
