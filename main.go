package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agungdwiprasetyo/api.agungdwiprasetyo.com/src/shared"
	"github.com/agungdwiprasetyo/go-utils"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		multiError := utils.NewMultiError()
		multiError.Append("authorization", fmt.Errorf("invalid signature"))

		response := shared.NewHTTPResponse(http.StatusUnauthorized, "failed", multiError)
		response.JSON(w)
	})

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
