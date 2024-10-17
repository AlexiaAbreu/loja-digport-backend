package routes

import (
	"net/http"

	"github.com/AlexiaAbreu/loja-digport-backend/controller"
	"github.com/gorilla/mux"
)

func HandleRequests() {
	route := mux.NewRouter()

	route.HandleFunc("/produtos", controller.BuscaProdutosHandler).Methods("GET")
	route.Handle("/produto", controller.AuthMiddleware(http.HandlerFunc(controller.BuscaProdutoPorNomeHandler))).Methods("GET")

	route.HandleFunc("/produto", controller.CriaProdutoHandler).Methods("POST")
	route.HandleFunc("/produto/{id}", controller.RemoveProdutoHandler).Methods("DELETE")
	route.HandleFunc("/produtos", controller.AtualizaProdutoHandler).Methods("PUT")

	//Usuario
	route.HandleFunc("/usuarios", controller.CriarUsuarioHandler).Methods("POST")
	route.HandleFunc("/usuarios", controller.BuscaUsuarioPorEmail).Methods("GET")
	route.HandleFunc("/usuarios/login", controller.LoginHandler).Methods("POST")

	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
	// 	AllowedHeaders:   []string{"Content-Type"},
	// 	AllowCredentials: true,
	// })

	// handler := c.Handler(route)

	http.ListenAndServe(":8080", route)

}
