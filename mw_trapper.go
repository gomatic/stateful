package main

import (
	"net/http"

	"log"
)

//
func trapper(f http.HandlerFunc) http.HandlerFunc {
	//
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %s: %+v", err, r)
			}
		}()
		f(w, r)
	}
}
