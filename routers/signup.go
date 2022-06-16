package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/vendenta/database"
	"github.com/vendenta/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	var u models.Account

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(u.Profile.FirstName) == 0 {
		http.Error(w, "debe llenar el campo nombre", 400)
		return
	}

	if len(u.Profile.LastName) == 0 {
		http.Error(w, "debe llenar el campo apellido", 400)
		return
	}

	if len(u.Password) < 6 {
		http.Error(w, "la contraseña debe ser mayor a 6 caracteres", 400)
		return
	}

	usuario := db.CreateUserProfile(u)

	usuario.Password = ""

	/*usuario, err := db.SearshProfile("2a3f423c-fe9d-44f0-97d2-ca1d143e5fa4")

	if err != nil {
		fmt.Println("usuario no encontrado")
	}*/

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(usuario)

}
