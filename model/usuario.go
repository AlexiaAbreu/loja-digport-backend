package model

import (
	"fmt"

	"github.com/AlexiaAbreu/loja-digport-backend/db"
	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func CriaUsuario(usuario Usuario) error {
	hash, err := hashPassword(usuario.Senha)
	if err != nil {
		return fmt.Errorf("erro ao criar usuário %s: %w", usuario.Nome, err)
	}
	db := db.ConectaBancoDados()

	_, err = db.Exec("INSERT INTO USUARIO VALUES ($1, $2, $3)", usuario.Nome, usuario.Email, usuario.Senha, hash)

	if err != nil {
		fmt.Errorf("erro ao inserir no banco de dados: %w", err)
	}

	return nil

}

func hashPassword(senha string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(senha), 14)
	if err != nil {
		return "", fmt.Errorf("erro ao tentar gerar hash da senha: %w", err)
	}
	return string(bytes), nil
}
