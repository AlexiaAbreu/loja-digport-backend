package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler) //1 parâmetro é o endpoint que irá retornar os dados para quem chamar
	http.ListenAndServe(":8080", nil)             //2 parâmetro é o handler em si, que pega a lista de produtos e transforma em JSON
}

func produtosHandler(writer http.ResponseWriter, reader *http.Request) {
	catalogoDeProdutos := retornaCatalogo()

	json.NewEncoder(writer).Encode(catalogoDeProdutos) // Aqui é feita a escrita do JSON a
}
