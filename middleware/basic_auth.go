package middleware

import (
	"fmt"
	"net/http"

	utils "github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/reverse-proxy/src/handler"
	"github.com/agungdwiprasetyo/reverse-proxy/src/shared"
)

// BasicAuth middleware
func BasicAuth(next http.HandlerFunc, username, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != username || pass != password {
			multiError := utils.NewMultiError()
			multiError.Append("authorization", fmt.Errorf("Invalid signature"))
			response := shared.NewHTTPResponse(http.StatusUnauthorized, "failed", multiError)
			response.JSON(w)
			return
		}

		if r.URL.Path != "/" {
			handler.NotFound(w, r)
			return
		}

		next(w, r)
	}
}
