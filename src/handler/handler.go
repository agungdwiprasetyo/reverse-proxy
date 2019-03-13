package handler

import (
	"fmt"
	"net/http"

	utils "github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/reverse-proxy/src/shared"
)

// Root handling root gateway request
func Root(w http.ResponseWriter, req *http.Request) {
	response := shared.NewHTTPResponse(http.StatusOK, "ok")
	response.JSON(w)
}

// NotFound handling request not found
func NotFound(w http.ResponseWriter, req *http.Request) {
	multiError := utils.NewMultiError()
	multiError.Append("router", fmt.Errorf("Resource '%s' not found", req.URL.Path))

	response := shared.NewHTTPResponse(http.StatusNotFound, "failed", multiError)
	response.JSON(w)
}
