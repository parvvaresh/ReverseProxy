# Reverse Proxy with Round-Robin Load Balancing in Go

This project implements a simple **reverse proxy** with **round-robin load balancing** in Go. It distributes incoming HTTP requests evenly across multiple backend servers.

---

## Features

* HTTP reverse proxy
* Round-robin load balancing between backend servers
* Built using Go's standard `net/http/httputil` package
* Easily extensible to add caching, rate limiting, and more

---

## Prerequisites

* Go installed (version 1.16+ recommended)
* Two or more backend HTTP servers running for testing (see example below)

---

## Setup and Usage

1. Clone or download this repository.

2. Start backend servers. For example, create two simple servers like:

   ```go
   // server.go
   package main

   import (
       "fmt"
       "log"
       "net/http"
   )

   func main() {
       http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
           fmt.Fprintf(w, "Hello from backend server at %s!", r.Host)
       })

       log.Println("Starting backend server on :8081")
       log.Fatal(http.ListenAndServe(":8081", nil))
   }
   ```

   Run this code twice with different ports (`:8081` and `:8082`), e.g.:

   ```bash
   go run server1.go  # runs on :8081
   go run server2.go  # runs on :8082
   ```

3. Run the reverse proxy:

   ```bash
   go run main.go
   ```

4. Access the proxy in your browser or via curl:

   ```
   http://localhost:8080/
   ```

   Requests will be forwarded in a round-robin manner between your backend servers.

---

## How It Works

* The proxy listens on port `8080`.
* Incoming requests are distributed to backend servers (`8081`, `8082`, etc.) one by one.
* Uses `httputil.NewSingleHostReverseProxy` to forward requests and responses.

---

## Possible Improvements

* Add caching layer
* Implement rate limiting
* Support HTTPS / TLS
* More detailed logging and error handling

---

## License

This project is licensed under the MIT License.
