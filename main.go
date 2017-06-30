package main

import (
	"net/http"
	"time"
)

//
func main() {

	required := require()

	mux := http.NewServeMux()
	mux.HandleFunc("/panic", trapper(panicking))
	mux.HandleFunc("/", trapper(required(ok)))

	s := &http.Server{
		Addr:           ":3000",
		Handler:        mux,
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
