package main

import "net/http"

func main() {
	http.HandleFunc("/", func(respWriter http.ResponseWriter, request *http.Request) {
		respWriter.WriteHeader(http.StatusOK)
		respWriter.Write([]byte("Hello, world!"))
	})
	http.ListenAndServe("0.0.0.0:80", nil)
}
