package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func main() {
    // 定义反向代理服务器的地址
    targetUrl, err := url.Parse("https://www.baidu.com/")
    if err != nil {
        log.Fatal(err)
    }

    // 创建反向代理服务器
    proxy := httputil.NewSingleHostReverseProxy(targetUrl)

    // 处理客户端请求
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // 设置Header
        w.Header().Set("X-Proxy", "go-reverse-proxy")

        // 转发请求
        proxy.ServeHTTP(w, r)
    })

    // 启动服务器
    fmt.Println("Starting server on :9090")
    if err := http.ListenAndServe(":9090", nil); err != nil {
        log.Fatal(err)
    }
}
