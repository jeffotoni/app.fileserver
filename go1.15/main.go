package main

import (
	"log"
	"net/http"

	_ "github.com/jeffotoni/app.fileserver/go1.15/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/",
		http.StripPrefix("/", http.FileServer(statikFS)))
	mux.HandleFunc("/api/ping", Ping)
	http.ListenAndServe(":8085", mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}
