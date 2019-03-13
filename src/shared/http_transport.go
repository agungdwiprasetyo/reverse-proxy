package shared

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/reverse-proxy/helper"
)

// Transport model
type Transport struct {
}

type monitoring struct {
	UserAgent string
	Method    string
	Path      string
	Duration  string
}

// RoundTrip is the ability to execute a single HTTP transaction
func (t *Transport) RoundTrip(request *http.Request) (*http.Response, error) {
	start := time.Now()

	response, err := http.DefaultTransport.RoundTrip(request)
	if err != nil {
		me := utils.NewMultiError()
		me.Append("server", fmt.Errorf("Server is not reachable. Server not working. Try again later"))

		respBody := httpResponse{
			Code:    http.StatusBadGateway,
			Message: "Bad Gateway",
			Errors:  me.ToMap(),
		}

		bg, _ := json.Marshal(respBody)
		response = &http.Response{
			Status:        "502 Bad Gateway",
			StatusCode:    http.StatusBadGateway,
			Proto:         "HTTP/1.1",
			ProtoMajor:    1,
			ProtoMinor:    1,
			Body:          ioutil.NopCloser(bytes.NewBuffer(bg)),
			ContentLength: int64(len(bg)),
			Request:       request,
			Header:        http.Header{"Content-Type": []string{"application/json"}},
		}
	}

	end := time.Now()

	statusColor := helper.ColorForStatus(response.StatusCode)
	methodColor := helper.ColorForMethod(request.Method)
	fmt.Fprintf(os.Stdout, "%s [PROXY] %s : %v | %s %3d %s | %13v | %15s | %s %-7s %s %s\n",
		helper.Magenta, helper.Reset,
		time.Now().Format("2006/01/02 - 15:04:05"),
		statusColor, response.StatusCode, helper.Reset,
		end.Sub(start),
		request.RemoteAddr,
		methodColor, request.Method, helper.Reset,
		request.RequestURI,
	)

	return response, nil
}
