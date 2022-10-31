package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mnarvaezm96/doom_go/db"
	"github.com/mnarvaezm96/doom_go/models"
)

/* Email valor usado en todos los endPoints */
var Email string

/* IDusuario devuelto del modelo, que se usuará en todos los endPoints*/
var IDusuario string

/*ProesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MiraTuEsaMonda")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Formato de Token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(Token *jwt.Token) (interface{}, error) {
		return miClave, nil

	})
	if err == nil {
		_, encontrado, _ := db.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDusuario = claims.ID.Hex()

		}
		return claims, encontrado, IDusuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err

}
