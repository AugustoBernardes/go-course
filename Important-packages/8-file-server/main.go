package main

import "net/http"

func main() {
	fileServer := http.FileServer(http.Dir("Important-packages/8-file-server/public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileServer)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello blog"))
	})
	http.ListenAndServe(":8080", mux)

}
