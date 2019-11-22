package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/agungdwiprasetyo/reverse-proxy/config"
	"github.com/agungdwiprasetyo/reverse-proxy/helper"
	"github.com/agungdwiprasetyo/reverse-proxy/middleware"
	"github.com/agungdwiprasetyo/reverse-proxy/src/handler"
	"github.com/agungdwiprasetyo/reverse-proxy/src/proxy"
)

func main() {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(helper.StringRed(err))
	}

	// init config
	conf := config.Init(appPath)

	// register all service proxy from config
	for _, service := range conf.Services {
		service.Root = fmt.Sprintf("/%s/", strings.Trim(service.Root, "/"))

		prx := proxy.NewProxy(service.Root, service.Host)
		http.HandleFunc(service.Root, prx.Handle)

		fmt.Fprintf(os.Stdout, "%s[GATEWAY]%s %s %s %s\n",
			helper.White, helper.Reset, helper.StringRed(fmt.Sprintf(":%d%s", conf.GatewayPort, service.Root)),
			helper.StringYellow("|===>"), helper.StringGreen(service.Host),
		)
	}

	// Handle root gateway
	root := middleware.BasicAuth(handler.Root, conf.Key.Username, conf.Key.Password)
	http.HandleFunc("/", middleware.Logger(root))

	port := fmt.Sprintf(":%d", conf.GatewayPort)
	log.Println("Server running on port", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(helper.StringRed(err))
	}
}
