package routes

import (
	"net/http"

	"github.com/AlexiaAbreu/loja-digport-backend/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/produto", controller.BuscaProdutoPorNomeHandler).Methods("GET")
	route.HandleFunc("/produto", controller.CriaProdutoHandler).Methods("POST")
	//route.HandleFunc("/produtos", controller.RemoveProdutoHandler).Methods("DELETE")
	//route.HandleFunc("/produtos", controller.AtualizaProdutoHandler).Methods("PUT")

	http.ListenAndServe(":8080", route)

}
