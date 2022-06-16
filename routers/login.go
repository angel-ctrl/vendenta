package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/vendenta/database"
	"github.com/vendenta/jwt"
	"github.com/vendenta/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var t models.Account

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario o contraseña invalido "+err.Error(), 400)
		return
	}

	if len(t.User) == 0 {
		http.Error(w, "el usuario es requerido ", 400)
		return
	}

	user, encontrado := db.SearchProfile(t.User)

	if !encontrado {
		http.Error(w, "usuario no encontrado ", 400)
		return
	}

	_, err = db.Loginatempt(user, t.Password)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalido "+err.Error(), 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(user)

	if err != nil {
		http.Error(w, "ocurrio un error "+err.Error(), 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(jwtKey)

}
