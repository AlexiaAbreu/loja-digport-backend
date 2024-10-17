package main

import (
	"github.com/AlexiaAbreu/loja-digport-backend/db"
	_ "github.com/lib/pq"
)

func main() {
	db.InitDB()
	StartServer()

	// produto := model.Produto{Nome: "meia", Preco: 17.99}
	// fmt.Println("preço", produto.Preco)

	// model.AumentaPreco(&produto)
	// fmt.Println("novo valor:", produto.Preco)

}
