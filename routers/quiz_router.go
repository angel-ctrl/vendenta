package routers

import (
	"encoding/json"
	"net/http"

	db "github.com/vendenta/database"
	"github.com/vendenta/models"
)

func QuizCreator(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		ID := r.URL.Query().Get("id")

		quiz := db.SearshQuizDB(ID)

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(quiz)

	case "POST":
		var u models.Quiz

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			http.Error(w, "error en los datos recibidos "+err.Error(), 400)
			return
		}

		quiz := db.CreateQuiz(u)

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(quiz)

	case "PUT":

		var u models.Quiz

		err := json.NewDecoder(r.Body).Decode(&u)

		if err != nil {
			http.Error(w, "error en los datos recibidos "+err.Error(), 400)
			return
		}

		quiz := db.Update_Quiz_Database(r.URL.Query().Get("id"), &u)

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(quiz)

	default:

		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return

	}
}


 