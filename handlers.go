package main

import "net/http"

func readyHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if path == "/" {
		http.FileServer(http.Dir("static")).ServeHTTP(w, r)
	} else {
		//http.FileServer(http.Dir("static/index.html")).ServeHTTP(w, r)
	}
}
