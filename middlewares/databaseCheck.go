package middlewares

import (
	"net/http"

	db "github.com/vendenta/database"
)

func CheckBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if db.PingDataBase() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}

		next.ServeHTTP(w, r)

	}

}
