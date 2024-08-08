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
	route.HandleFunc("/produtos", controller.AdicionaProdutoHandler).Methods("POST")

	http.ListenAndServe(":8080", route)

}
