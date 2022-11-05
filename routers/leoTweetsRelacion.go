package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mnarvaezm96/doom_go/db"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := db.LeoTweetSeguidores(IDusuario, pagina)
	if correcto == false {
		http.Error(w, "Error a leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Contente-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
