package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")
	var extesion = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDusuario + "." + extesion
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDusuario + "." + extesion
	status, err = db.ModificoResgistro(usuario, IDusuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la DB"+err.Error(), http.StatusBadRequest)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
