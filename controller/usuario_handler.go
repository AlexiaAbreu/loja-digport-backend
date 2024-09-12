package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AlexiaAbreu/loja-digport-backend/model"
)

func CriaUsuarioHandler(w http.ResponseWriter, r *http.Request) {
	var usuario model.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	error := model.CriaUsuario(usuario)
	if error != nil {
		w.WriteHeader(http.StatusBadGateway)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
