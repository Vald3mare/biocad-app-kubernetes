package main

import (
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<h1>Hello, world!</h1>"))
}

func main() {
	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":32777", nil)
	if err != nil {
		panic(err)
	}
}
