package main

import (
	"net/http"

	"log"
)

//
func trapper(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
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
