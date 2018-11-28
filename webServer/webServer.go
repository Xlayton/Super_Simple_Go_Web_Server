package main

import (
	"net/http"
)

func main() {
	fsHTML := http.FileServer(http.Dir("./public"))
	fsCSS := http.FileServer(http.Dir("./public/css"))
	fsJS := http.FileServer(http.Dir("./public/js"))
	http.Handle("/public/js", http.StripPrefix("/public/js", fsJS))
	http.Handle("/public/css", http.StripPrefix("/public/css", fsCSS))
	http.Handle("/", http.StripPrefix("/", fsHTML))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
		return
	}
}
