package main

import (
	"github.com/AlexiaAbreu/loja-digport-backend/model"
)

var ListaDeProdutos []model.Product = []model.Product{}

func retornaProdutoPeloNome(nomeDoProduto string) []model.Product {
	listaComProdutosFiltrados := []model.Product{}

	for _, produto := range ListaDeProdutos {
		if produto.Name == nomeDoProduto {
			listaComProdutosFiltrados = append(listaComProdutosFiltrados, produto)
		}
	}

	return listaComProdutosFiltrados
}

func registraProduto(novoProduto model.Product) {
	ListaDeProdutos = append(ListaDeProdutos, novoProduto)
}
