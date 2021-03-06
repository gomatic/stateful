package main

import (
	"fmt"
	"log"
	"net/http"
)

//
func ok(w http.ResponseWriter, r *http.Request) {
	switch path := r.URL.Path[1:]; path {
	case "error":
		switch ctx := r.Context().(type) {
		case *svctx:
			log.Println("failing")
			ctx.initializeFailure(fmt.Errorf("failed"))
		}
	default:
		ctx, ok := r.Context().(*svctx)
		if !ok {
			log.Printf("Unexpected context type: %T", r.Context())
			return
		}
		db, acc := ctx.db, ctx.acc
		fmt.Fprintf(w, "handling %s db:%+v acc:%+v\n", path, *db, *acc)
	}
}
