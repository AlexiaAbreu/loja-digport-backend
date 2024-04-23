package main

import (
	"github.com/AlexiaAbreu/loja-digport-backend/model"
)

func retornaTodoEstoque() []model.Product {
	estoque := []model.Product{
		{
			ID:          "123",
			Type:        "roupa",
			Name:        "Blusa Social",
			Description: "Blusa Social preta",
			Price:       100,
			Amount:      5,
			Image:       "https://images.tcdn.com.br/img/img_prod/766954/calca_social_feminina_sem_bolso_113_1_20200410203252.jpg",
		},
		{
			ID:          "234",
			Type:        "roupa",
			Name:        "Calça Social",
			Description: "Calça Social Preta",
			Price:       200,
			Amount:      10,
			Image:       "https://tfcpr5.vtexassets.com/arquivos/ids/156896-800-auto?v=638180210839570000&width=800&height=auto&aspect=true",
		},
		{
			ID:          "345",
			Type:        "Acessório",
			Name:        "Cinto",
			Description: "Cinto preto",
			Price:       20,
			Amount:      10,
			Image:       "https://cdn.awsli.com.br/600x450/1944/1944028/produto/237999112/img_0094-lx2njfzyca.JPG",
		},
	}
	//estoque := make([]model.Product, len(catalogo))

	return estoque
}

func retornaProdutoPeloNome(nomeDoProduto string) model.Product {
	var produto model.Product

	if nomeDoProduto == "Blusa Social" {
		produto = model.Product{
			ID:          "123",
			Type:        "roupa",
			Name:        "Blusa Social",
			Description: "Blusa Social preta",
			Price:       100,
			Amount:      5,
			Image:       "https://images.tcdn.com.br/img/img_prod/766954/calca_social_feminina_sem_bolso_113_1_20200410203252.jpg",
		}
	}
	if nomeDoProduto == "Calça Social" {
		produto = model.Product{
			ID:          "234",
			Type:        "roupa",
			Name:        "Calça Social",
			Description: "Calça Social Preta",
			Price:       200,
			Amount:      10,
			Image:       "https://tfcpr5.vtexassets.com/arquivos/ids/156896-800-auto?v=638180210839570000&width=800&height=auto&aspect=true",
		}
	}
	if nomeDoProduto == "Acessório" {
		produto = model.Product{
			ID:          "345",
			Type:        "Acessório",
			Name:        "Cinto",
			Description: "Cinto preto",
			Price:       20,
			Amount:      10,
			Image:       "https://cdn.awsli.com.br/600x450/1944/1944028/produto/237999112/img_0094-lx2njfzyca.JPG",
		}
	}

	return produto
}
