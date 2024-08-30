package model

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/AlexiaAbreu/loja-digport-backend/db"
	"github.com/google/uuid"
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
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}
	produtos := []Produto{}

	for resultado.Next() {
		err = resultado.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
		if err != nil {
			panic(err.Error())
		}
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

func BuscaProdutoPorNome(nomeDoProduto string) Produto {
	db := db.ConectaBancoDados()

	res := db.QueryRow("SELECT * FROM produtos WHERE nome = $1", nomeDoProduto)

	err := res.Scan(&id, &nome, &preco, &descricao, &imagem, &quantidade)
	if err == sql.ErrNoRows {
		fmt.Printf("Produto não encontrado %s\n", nome)
	} else if err != nil {
		panic(err.Error())
	}

	var p Produto
	p.ID = id
	p.Nome = nome
	p.Descricao = descricao
	p.Preco = preco
	p.Imagem = imagem
	p.Quantidade = quantidade

	defer db.Close()
	return p
}

func CriaProduto(prod Produto) error {

	if produtoCadastrado(prod.Nome) {
		fmt.Printf("Produto já cadastrado %s\n", prod.Nome)
		return fmt.Errorf("Produto já cadastrado")
	}

	db := db.ConectaBancoDados()
	id := uuid.NewString()
	nome := prod.Nome
	preco := prod.Preco
	descricao := prod.Descricao
	imagem := prod.Imagem
	quantidade := prod.Quantidade

	strInsert := "INSERT INTO produtos VALUES($1, $2, $3, $4, $5, $6)"

	result, err := db.Exec(strInsert, id, nome, strconv.FormatFloat(preco, 'f', 1, 64), descricao,
		imagem, strconv.Itoa(quantidade))
	if err != nil {
		panic(err.Error())
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Produto %s criado com sucesso (%d row affected)", id, rowsAffected)

	defer db.Close()
	return nil
}

func produtoCadastrado(nomeDoProduto string) bool {
	prod := BuscaProdutoPorNome(nomeDoProduto)

	return prod.Nome == nomeDoProduto
}

func RemoveProduto(id string) error {
	db := db.ConectaBancoDados()
	defer db.Close()

	_, err := db.Exec("DELETE FROM PRODUTOS WHERE id = $1", id)

	if err != nil {

		fmt.Println("erro ao deletar")
		return fmt.Errorf("Erro: %w", err)
	}
	fmt.Println("Sucesso ao deletar")
	return nil

}

// func UpdateProduto(produto Produto) error {

// }
