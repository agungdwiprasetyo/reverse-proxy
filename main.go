package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/agungdwiprasetyo/api.agungdwiprasetyo.com/src/proxy"
	"github.com/agungdwiprasetyo/api.agungdwiprasetyo.com/src/shared"
	utils "github.com/agungdwiprasetyo/go-utils"
	"github.com/agungdwiprasetyo/go-utils/debug"
)

// Config app
type Config struct {
	GatewayPort string `json:"gatewayPort"`
	Proxy       []struct {
		Root string `json:"root"`
		Host string `json:"host"`
	} `json:"proxy"`
}

func main() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/config.json", appPath))
	if err != nil {
		log.Fatal(err)
	}

	var config Config
	if err := json.Unmarshal(b, &config); err != nil {
		log.Fatal(err)
	}

	for _, pr := range config.Proxy {
		prx := proxy.NewProxy(pr.Root, pr.Host)
		http.HandleFunc(pr.Root, prx.Handle)
		debug.Println(pr.Root, pr.Host)
	}

	// Handle root gateway
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		multiError := utils.NewMultiError()
		multiError.Append("authorization", fmt.Errorf("invalid signature"))

		response := shared.NewHTTPResponse(http.StatusUnauthorized, "failed", multiError)
		response.JSON(w)
	})

	port := fmt.Sprintf(":%s", config.GatewayPort)
	log.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
