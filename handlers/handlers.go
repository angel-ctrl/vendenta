package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/vendenta/middlewares"
	"github.com/vendenta/routers"
	"github.com/vendenta/webSockets"
	websocketstestpause "github.com/vendenta/webSocketsTestPause"
)

func HandRouters() {

	router := mux.NewRouter()

	router.HandleFunc("/API/register", middlewares.CheckBD(routers.Register)).Methods("POST")
	router.HandleFunc("/API/login", middlewares.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/API/user_example", middlewares.CheckBD(routers.UserExampleCreator)).Methods("POST")
	router.HandleFunc("/API/ws", middlewares.CorsHeaders(middlewares.CheckBD(webSockets.HandlerWebSocket))).Methods("GET")
	router.HandleFunc("/API/wsDisplay", websocketstestpause.HandlerWebSocketPause).Methods("GET")
	router.HandleFunc("/API/Quiz", middlewares.CorsHeaders(middlewares.CheckBD(middlewares.Valid_token(routers.QuizCreator)))).Methods("GET", "POST", "PUT")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	cor := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://25.12.136.60:3000", "http://127.0.0.1:8000"},
		AllowedHeaders:   []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token", "Authorization", "Accept", "Accept-Language"},
		AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE"},
		Debug:            true,
		AllowCredentials: true,
	})

	handler := cor.Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
