package main

import "net/http"

func main() {
	http.HandleFunc("/first-run", func(rw http.ResponseWriter, r *http.Request) {

		rw.Write([]byte("Hello World"))

	})
	http.ListenAndServe(":8080", nil)
}
