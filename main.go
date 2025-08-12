package main

import (
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
    "sync"
    "time"
)

type Proxy struct {
    target *url.URL
    proxy  *httputil.ReverseProxy
}

func NewProxy(target string) *Proxy {
    url, err := url.Parse(target)
    if err != nil {
        log.Fatalf("Error parsing target URL: %v", err)
    }
    return &Proxy{
        target: url,
        proxy:  httputil.NewSingleHostReverseProxy(url),
    }
}

func main() {
    targets := []string{
        "http://localhost:8081",
        "http://localhost:8082",
    }

    proxies := []*Proxy{}
    for _, t := range targets {
        proxies = append(proxies, NewProxy(t))
    }

    var mu sync.Mutex
    var current int

    handler := func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        proxy := proxies[current]
        current = (current + 1) % len(proxies)
        mu.Unlock()

        log.Printf("Proxying request to %s%s", proxy.target.Host, r.URL.Path)


        proxy.proxy.ServeHTTP(w, r)
    }

    srv := &http.Server{
        Addr:         ":8080",
        Handler:      http.HandlerFunc(handler),
        ReadTimeout:  10 * time.Second,
        WriteTimeout: 10 * time.Second,
    }

    log.Println("Reverse proxy started on :8080")
    log.Fatal(srv.ListenAndServe())
}
