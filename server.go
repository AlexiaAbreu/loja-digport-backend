package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	roteador := mux.NewRouter()
	roteador.HandleFunc("/produtos", produtosHandler).Methods("GET")
	roteador.HandleFunc("/produtos/{nomeDoProduto}", produtosPorNome).Methods("GET")

	http.ListenAndServe(":8080", roteador)
}

func produtosHandler(writer http.ResponseWriter, request *http.Request) {
	catalogoDeProdutos := retornaTodoEstoque()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(catalogoDeProdutos)
}

func produtosPorNome(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	nomeDoProduto := params["nomeDoProduto"]

	catalogoDeProdutosEspecificos := retornaProdutoPeloNome(nomeDoProduto)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(catalogoDeProdutosEspecificos)
}
