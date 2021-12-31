package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/jeffotoni/app.fileserver/go1.15/statik"
	"github.com/rakyll/statik/fs"
)

type LoginPage struct {
	IfLabelone string
	Labelone   string
	MsgError   string
}

func LoginHtml(w http.ResponseWriter, req *http.Request) {
	tmpl := template.Must(template.ParseFiles("./web/index.html"))
	data := LoginPage{
		Labelone: "Sign in",
	}
	tmpl.Execute(w, data)
}

func main() {
	statikFS, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHtml)
	mux.Handle("/",
		http.StripPrefix("/", http.FileServer(statikFS)))
	mux.HandleFunc("/api/ping", Ping)
  log.Println("Run server :8085")
	http.ListenAndServe(":8085", mux)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte(`pong`))
}
