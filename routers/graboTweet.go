package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDusuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := db.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logardo intesertar el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
