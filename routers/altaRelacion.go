package routers

import (
	"net/http"

	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDusuario
	t.UsurioRelacionID = ID

	status, err := db.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relación"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar la relación"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
