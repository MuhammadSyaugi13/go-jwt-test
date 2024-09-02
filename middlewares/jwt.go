package middlewares

import (
	"go-jwt-web/config"
	"go-jwt-web/helper"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {

			//jika tidak ada cookie sama sekali
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unautorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		//mengambil token value
		tokenString := c.Value

		claims := &config.JWTClaim{}

		// parsing token jwt
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		//validasi token jwt
		if err != nil {
			v, _ := err.(*jwt.ValidationError)

			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{"message": "Signature Invalid"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return

			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "Unautorized, Token Expired"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "Unautorizedd"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
		}

		//jika token tidak valid
		if !token.Valid {
			response := map[string]string{"message": "Token Invalid"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		// jika berhasil melewati semua filter, maka lanjutkan ke halaman yang dituju
		next.ServeHTTP(w, r)

	})
}
