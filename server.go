package main

import (
	"encoding/json"
	"net/http"

	"github.com/AlexiaAbreu/loja-digport-backend/model"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getProdutos(w, r)
	} else if r.Method == "POST" {
		addProdutos(w, r)
	}
}

func getProdutos(w http.ResponseWriter, r *http.Request) {
	queryNome := r.URL.Query().Get("name")
	if queryNome != "" {
		produtosFiltradosPeloNome := retornaProdutoPeloNome(queryNome)
		json.NewEncoder(w).Encode(produtosFiltradosPeloNome)
	} else {
		produtos := ListaDeProdutos
		json.NewEncoder(w).Encode(produtos)
	}
}

func addProdutos(w http.ResponseWriter, r *http.Request) {
	var produto model.Product
	json.NewDecoder(r.Body).Decode(&produto)

	registraProduto(produto)

	w.WriteHeader(http.StatusCreated)
}
