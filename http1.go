package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World! \n")
}

func main() {
	//http.HandleFunc("/", hello)
	fs := http.FileServer(http.Dir("/"))
	http.Handle("/", fs)
	http.ListenAndServe(":1337", nil)
}
