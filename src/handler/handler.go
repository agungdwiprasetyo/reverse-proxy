package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	utils "github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/reverse-proxy/config"
	"github.com/agungdwiprasetyo/reverse-proxy/src/shared"
)

// Root handling root gateway request
func Root(w http.ResponseWriter, req *http.Request) {
	var response = map[string]interface{}{
		"message":  "ok",
		"services": config.GlobalConfig.Services,
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Origin", "Agung Dwi Prasetyo")
	w.Header().Set("Go-Version", runtime.Version())
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}

// NotFound handling request not found
func NotFound(w http.ResponseWriter, req *http.Request) {
	multiError := utils.NewMultiError()
	multiError.Append("router", fmt.Errorf("Resource '%s' not found", req.URL.Path))

	response := shared.NewHTTPResponse(http.StatusNotFound, "failed", multiError)
	response.JSON(w)
}
