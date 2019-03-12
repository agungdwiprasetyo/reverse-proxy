package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/agungdwiprasetyo/api.agungdwiprasetyo.com/src/shared"
)

type Proxy struct {
	target *url.URL
	root   string
	proxy  *httputil.ReverseProxy
}

func NewProxy(root, target string) *Proxy {
	url, _ := url.Parse(target)

	return &Proxy{target: url, root: root, proxy: httputil.NewSingleHostReverseProxy(url)}
}

func (l *Proxy) Handle(w http.ResponseWriter, r *http.Request) {
	i := strings.Index(r.URL.Path, l.root)
	if i >= 0 {
		r.URL.Path = r.URL.Path[i+len(l.root):]
	}

	l.proxy.Transport = new(shared.Transport)
	l.proxy.ServeHTTP(w, r)

}
