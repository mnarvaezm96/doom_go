package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mnarvaezm96/doom_go/middlew"
	"github.com/mnarvaezm96/doom_go/routers"
	"github.com/rs/cors"
)

/* Manejadores seteo mi puerto , el handler y pongo a escuchar el server*/
func Manejadores() {

	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	/* router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Verperfil))).Methods("GET") */

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
