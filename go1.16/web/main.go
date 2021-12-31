package main

import (
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	HTTP_PORT = "0.0.0.0:8080"
)

//go:embed static/* static/css/* static/fonts/* static/images/* static/js/*
//go:embed static/index.html
var contentfs embed.FS

type LoginPage struct {
	IfLabelone string
	Labelone   string
	MsgError   string
}

func LoginHtml(w http.ResponseWriter, req *http.Request) {
	//tmpl := template.Must(template.ParseFiles("./static/index.html"))
	//html, err := contentfs.ReadFile("./static/index.html")
	//if err != nil {
	//	log.Println("error:", err)
  //	return
	//}
	tmpl, err := template.ParseFS(contentfs, "static/index.html")
	if err!=nil{
		log.Println("error:",err)
		return 
	}
	//tmpl, err := template.New("web").Parse(string(html))
	data := LoginPage{
		Labelone: "Sign in",
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", LoginHtml)
	//mux.Handle("/", http.StripPrefix("/", DisabledFs(fs)))
	fs := http.FileServer(http.FS(contentfs))
	mux.Handle("/", http.StripPrefix("/static/", DisabledFs(fs)))
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
