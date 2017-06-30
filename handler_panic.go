package main

import (
	"net/http"
)

//
func panicking(http.ResponseWriter, *http.Request) {
	panic("trap me")
}
