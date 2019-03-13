package helper

import (
	"fmt"
	"net/http"
)

var (
	Green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	White   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	Yellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	Red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	Blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	Magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	Cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	Reset   = string([]byte{27, 91, 48, 109})
)

// StringRed return str with red color
func StringRed(str interface{}) string {
	return fmt.Sprintf("\x1b[31;1m%+v\x1b[0m", str)
}

// StringYellow return str with yellow color
func StringYellow(str interface{}) string {
	return fmt.Sprintf("\x1b[33;1m%+v\x1b[0m", str)
}

// StringGreen return str with green color
func StringGreen(str interface{}) string {
	return fmt.Sprintf("\x1b[32;1m%+v\x1b[0m", str)
}

// ColorForStatus func
func ColorForStatus(code int) string {
	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return Green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return White
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return Yellow
	default:
		return Red
	}
}

// ColorForMethod func
func ColorForMethod(method string) string {
	switch method {
	case http.MethodGet:
		return Blue
	case http.MethodPost:
		return Cyan
	case http.MethodPut:
		return Yellow
	case http.MethodDelete:
		return Red
	case http.MethodPatch:
		return Green
	case http.MethodHead:
		return Magenta
	case http.MethodOptions:
		return White
	default:
		return Reset
	}
}
