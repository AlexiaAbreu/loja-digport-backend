package routes

import (
	"net/http"

	"github.com/AlexiaAbreu/loja-digport-backend/tree/main/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/procutos", controller.BuscaProdutosHandler).Methods("GET")
	route.HandleFunc("/procuto", controller.BuscaProdutosPorNomeHandler).Methods("GET")
	route.HandleFunc("/procutos", controller.CriaProdutosHandler).Methods("POST")

	http.ListenAndServe(":8080", route)

}
