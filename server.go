package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler)  //1 parâmetro é o endpoint que irá retornar os dados para quem chamar, 2 parâmetro é o handler em si, que pega a lista de produtos e transforma em JSON
	http.HandleFunc("/produtos/", produtosPorNome) // ??

	http.ListenAndServe(":8080", nil)
}

func produtosHandler(writer http.ResponseWriter, reader *http.Request) {
	catalogoDeProdutos := retornaTodoEstoque()

	json.NewEncoder(writer).Encode(catalogoDeProdutos) // Aqui é feita a escrita do JSON a
}

func produtosPorNome(writer http.ResponseWriter, reader *http.Request) {
	catalogoDeProdutos := retornaProdutoPeloNome()

	json.NewEncoder(writer).Encode(catalogoDeProdutos)
}
