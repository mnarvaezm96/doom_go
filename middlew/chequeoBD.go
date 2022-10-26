package middlew

import (
	"net/http"

	"github.com/mnarvaezm96/doom_go/db"
)

/* ChequeoBD es el middlew que me permite conocer el estado de la base de datos */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if db.ChequeoConnection() == 0 {

			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		Next.ServeHTTP(w, r)
	}

}
