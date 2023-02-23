package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/jefrydep/go_backend/middlew"
	"github.com/jefrydep/go_backend/routers"
)

// seteando el puerto
func Manejadores() {
	router := mux.NewRouter()
	// rutas
	router.HandleFunc("/registro",middlew.CheckeoBD(routers.registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
