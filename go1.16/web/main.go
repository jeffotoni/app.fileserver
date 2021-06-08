package main

import (
	"embed"
	"log"
	"net/http"
	"strings"
)

var (
	HTTP_PORT = "0.0.0.0:8080"
)

//go:embed static
var contentfs embed.FS

func main() {
	mux := http.NewServeMux()
	//mux.Handle("/", http.StripPrefix("/", DisabledFs(fs)))
	fs := http.FileServer(http.FS(contentfs))
	mux.Handle("/", http.StripPrefix("/static", DisabledFs(fs)))
	mux.HandleFunc("/ping", Ping)

	log.Println("Run Server:", HTTP_PORT)
	http.ListenAndServe(HTTP_PORT, mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}

func DisabledFs(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/static") {
			http.NotFound(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
