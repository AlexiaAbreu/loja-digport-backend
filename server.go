package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/produtos", produtosHandler).Methods("GET")
	r.HandleFunc("/produtos/{nomeDoProduto}", produtosPorNome).Methods("GET")

	http.ListenAndServe(":8080", r)

	// http.ListenAndServe(":8080", r)
	// http.HandleFunc("/produtos", produtosHandler)  //1 parâmetro é o endpoint que irá retornar os dados para quem chamar, 2 parâmetro é o handler em si, que pega a lista de produtos e transforma em JSON
	// http.HandleFunc("/produtos/{nomeDoProduto}", produtosPorNome)
	// http.ListenAndServe(":8080", nil)
}

func produtosHandler(writer http.ResponseWriter, request *http.Request) {
	catalogoDeProdutos := retornaTodoEstoque()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(catalogoDeProdutos)
	// json.NewEncoder(writer).Encode(catalogoDeProdutos) // Aqui é feita a escrita do JSON a
}

func produtosPorNome(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	nomeDoProduto := params["nomeDoProduto"]

	catalogoDeProdutosEspecificos := retornaProdutoPeloNome(nomeDoProduto)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(catalogoDeProdutosEspecificos)
}
