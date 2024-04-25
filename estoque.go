package main

import (
	"errors"

	"github.com/AlexiaAbreu/loja-digport-backend/model"
)

var ListaDeProdutos []model.Produto = []model.Produto{}

func retornaProdutoPeloNome(nomeDoProduto string) []model.Produto {
	listaComProdutosFiltrados := []model.Produto{}

	for _, produto := range ListaDeProdutos {
		if produto.Nome == nomeDoProduto {
			listaComProdutosFiltrados = append(listaComProdutosFiltrados, produto)
		}
	}

	return listaComProdutosFiltrados
}

func adicionaNovoProduto(novoProduto model.Produto) error {
	for _, produto := range ListaDeProdutos {
		if novoProduto.ID == produto.ID {
			return errors.New("não é possível criar novo produto. ID já existente")
		}
	}
	ListaDeProdutos = append(ListaDeProdutos, novoProduto)
	return nil // equivalente ao null do go
}
