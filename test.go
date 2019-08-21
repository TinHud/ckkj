package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("./public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", index)
	mux.HandleFunc("/index", index)
	mux.HandleFunc("/datum", datum)
	mux.HandleFunc("/xiehui", xiehui)
	mux.HandleFunc("/project", project)
	server := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}

	server.ListenAndServe()

}

func xiehui(w http.ResponseWriter, r *http.Request) {
	bs, _ := ioutil.ReadFile("xiehui.html")
	w.Write(bs)
}

func project(w http.ResponseWriter, r *http.Request) {
	bs, _ := ioutil.ReadFile("project.html")
	w.Write(bs)
}

func datum(w http.ResponseWriter, r *http.Request) {
	bs, _ := ioutil.ReadFile("datum.html")
	w.Write(bs)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" && r.URL.Path != "/index" {
		bs, _ := ioutil.ReadFile("404.html")
		w.Write(bs)
		return
	}
	bs, _ := ioutil.ReadFile("index.html")
	w.Write(bs)
}
