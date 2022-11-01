package routers

import (
	"encoding/json"
	"net/http"

	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}
	var status bool

	status, err = db.ModificoResgistro(t, IDusuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente"+err.Error(), 400)
		return

	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
