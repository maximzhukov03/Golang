package reverse

import (
    "fmt"
    "net/http"
    "net/http/httputil"
    "net/url"
    "strings"
)

type ReverseProxy struct {
    host string
    port string
}

func NewReverseProxy(host, port string) *ReverseProxy {
    return &ReverseProxy{host: host, port: port}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        if strings.HasPrefix(path, "/api") ||
            strings.HasPrefix(path, "/swagger") ||
            strings.HasPrefix(path, "/mycustompath/pprof") {
            next.ServeHTTP(w, r)
            return
        }

        target := fmt.Sprintf("http://%s:%s", rp.host, rp.port)
        uri, err := url.Parse(target)
        if err != nil {
            http.Error(w, "Invalid proxy target", http.StatusInternalServerError)
            return
        }
        r.Header.Set("Reverse-Proxy", "true")

        proxy := httputil.ReverseProxy{
            Director: func(req *http.Request) {
                req.URL.Scheme = uri.Scheme
                req.URL.Host = uri.Host
                req.URL.Path = uri.Path + req.URL.Path
                req.Host = uri.Host
            },
        }
        proxy.ServeHTTP(w, r)
    })
}