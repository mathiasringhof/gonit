package main

import (
	"encoding/json"
	"net/http"
)

func setupHandlers() {
	auth := SimpleHeaderAuthenticator{"X-auth", "letmein"}
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api", auth.Wrap(apiHandler))
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir(config.WebDir))))
}

func indexHandler(rw http.ResponseWriter, req *http.Request) {
	http.Redirect(rw, req, "/web/index.html", http.StatusFound)
}

func apiHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		apiGetHandler(rw, req)
	case "PUT":
		apiPutHandler(rw, req)
	default:
		http.Error(rw, "Only GET & PUT allowed.", 405)
	}
}

func apiPutHandler(rw http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var r Repository
	if err := dec.Decode(&r); err != nil {
		http.Error(rw, "Bad request: "+err.Error(), 400)
	}
	if r.Exists() {
		http.Error(rw, "Repository "+r.Name+" already exists.", 409)
	}
	if err := r.Create(); err != nil {
		http.Error(rw, "Internal server error: "+err.Error(), 500)
	}
}

func apiGetHandler(rw http.ResponseWriter, req *http.Request) {
	enc := json.NewEncoder(rw)
	r := getRepositories()
	if err := enc.Encode(&r); err != nil {
		http.Error(rw, "Error encoding repositories to JSON", 500)
	}
}
