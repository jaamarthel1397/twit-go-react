package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jaamarthel1397/twit-go-react/bd"
	"github.com/jaamarthel1397/twit-go-react/models"
)

/* Email es el valor de email usado */
var Email string

/* IDUsuario es el ID devuelto del modelo, que se usara en todos los EndPoints */
var IDUsuario string

/* ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Viva el leon")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}

	return claims, false, string(""), err
}
