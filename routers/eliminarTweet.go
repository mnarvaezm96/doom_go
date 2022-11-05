package routers

import (
	"net/http"

	"github.com/mnarvaezm96/doom_go/db"
)

func EliminarTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el paremetro ID", http.StatusBadRequest)
		return
	}

	err := db.BorroTweet(ID, IDusuario)
	if err != nil {
		http.Error(w, "Ocurrio un errro al intentar borrar el tweet", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
