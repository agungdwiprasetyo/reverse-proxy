package shared

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"reflect"

	utils "github.com/agungdwiprasetyo/go-utils"
)

// HTTPResponse abstract interface
type HTTPResponse interface {
	JSON(w http.ResponseWriter)
	XML(w http.ResponseWriter)
}

type (
	// httpResponse model
	httpResponse struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Meta    *Meta       `json:"meta,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Errors  interface{} `json:"errors,omitempty"`
	}

	// Meta model
	Meta struct {
		Page         int `json:"page"`
		Limit        int `json:"limit"`
		TotalRecords int `json:"totalRecords"`
	}
)

// NewHTTPResponse for create common response, data must in first params and meta in second params
func NewHTTPResponse(code int, message string, params ...interface{}) HTTPResponse {
	commonResponse := new(httpResponse)

	for _, param := range params {
		refValue := reflect.ValueOf(param)
		if refValue.Kind() == reflect.Ptr {
			refValue = refValue.Elem()
		}
		param = refValue.Interface()

		switch data := param.(type) {
		case Meta:
			commonResponse.Meta = &data
		case utils.MultiError:
			commonResponse.Errors = data.ToMap()
		default:
			commonResponse.Data = param
		}
	}

	if code < 400 {
		commonResponse.Success = true
	}

	commonResponse.Code = code
	commonResponse.Message = message
	return commonResponse
}

// JSON for set http JSON response (Content-Type: application/json)
func (resp *httpResponse) JSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Origin", "Agung Dwi Prasetyo")
	w.WriteHeader(resp.Code)
	json.NewEncoder(w).Encode(resp)
}

// XML for set http XML response (Content-Type: application/xml)
func (resp *httpResponse) XML(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/xml")
	w.Header().Set("Origin", "Agung Dwi Prasetyo")
	w.WriteHeader(resp.Code)
	xml.NewEncoder(w).Encode(resp)
}
