package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello from server on port %s!", r.Host)
    })

    log.Println("Starting server on :8082")
    log.Fatal(http.ListenAndServe(":8082", nil))
}
