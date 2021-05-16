package main

import "net/http"

func main() {
	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {

	})
	http.ListenAndServe(":80", nil)
}
