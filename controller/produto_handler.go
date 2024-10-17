package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexiaAbreu/loja-digport-backend/model"
	"github.com/gorilla/mux"
)

func BuscaProdutosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := model.BuscaTodosProdutos()
	json.NewEncoder(w).Encode(produtos)

}

func BuscaProdutoPorNomeHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	produto := model.BuscaProdutoPorNome(nome)
	json.NewEncoder(w).Encode(produto)
}

func CriaProdutoHandler(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)

	error := model.CriaProduto(produto)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func RemoveProdutoHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	error := model.RemoveProduto(id)
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

func AtualizaProdutoHandler(w http.ResponseWriter, r *http.Request) {
	var produto model.Produto
	json.NewDecoder(r.Body).Decode(&produto)
	error := model.UpdateProduto(produto)
	if error != nil {
		fmt.Print(error)
		w.WriteHeader(http.StatusNotFound)
	} else {
		fmt.Println(produto.ID)
		w.WriteHeader(http.StatusOK)
	}

}
