package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)
	// Because struct is attached , we can call the function
	mux.Handle("/blob", Blob{title: "My Blob"})
	go http.ListenAndServe(":8080", mux)

	// Use if need multiple servers
	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Other server"))
	})
	http.ListenAndServe(":3000", mux2)
}

func HomeHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Hello world!"))
}

type Blob struct {
	title string
}

func (blob Blob) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte(blob.title))
}
