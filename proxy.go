package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    targetUrl, err := url.Parse("https://www.baidu.com/")
    if err != nil {
        log.Fatal(err)
    }

    proxy := httputil.NewSingleHostReverseProxy(targetUrl)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-Proxy", "go-reverse-proxy")
        r.Host = "www.baidu.com"
        r.URL.Scheme = "https"
        r.URL.Host = "www.baidu.com"
        proxy.ServeHTTP(w, r)
    })

    fmt.Println("Starting server on :9090")
    if err := http.ListenAndServe(":9090", nil); err != nil {
        log.Fatal(err)
    }
}
