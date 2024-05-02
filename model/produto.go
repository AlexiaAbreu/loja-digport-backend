package model

import (
	"github.com/AlexiaAbreu/loja-digport-backend/db"
)

type Produto struct {
	ID         string  `json:"id"`
	Nome       string  `json:"nome"`
	Preco      float64 `json:"preco"`
	Descricao  string  `json:"descricao"`
	Imagem     string  `json:"imagem"`
	Quantidade int     `json:"quantidadeEmEstoque"`
}

var id, nome string
var preco float64
var descricao, imagem string
var quantidade int

func BuscaTodosProdutos() []Produto {
	db := db.ConectaBancoDados()

	resultado, err := db.Query("SELECT * FROM produtos")
	err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	produto := Produto{}
	produtos := []Produto{}

	if err != nil {
		panic(err.Error())
	}

	for resultado.Next() {
		produto.ID = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Imagem = imagem
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}
	defer db.Close()
	return produtos
}
