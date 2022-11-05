package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/mnarvaezm96/doom_go/db"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := db.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	OPenFile, err := os.Open("uploads/banners" + perfil.Banner)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OPenFile)
	if err != nil {
		http.Error(w, "Error al copiar la iamgen", http.StatusBadRequest)
	}

}
