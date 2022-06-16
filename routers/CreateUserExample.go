package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/vendenta/database"
	"github.com/vendenta/models"
)

func UserExampleCreator(w http.ResponseWriter, r *http.Request) {

	var u models.UserExample

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "error en los datos recibidos "+err.Error(), 400)
		return
	}

	quiz := db.CreateUserExample(u)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(quiz)

}
