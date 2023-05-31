package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/greeter/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		name := parts[len(parts)-1]
		if name != "" {
			fmt.Fprintf(w, "Hello, %s!", name)
		} else {
			fmt.Fprint(w, "Hello, world!")
		}
	})

	http.ListenAndServe(":8080", nil)
}
