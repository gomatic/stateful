package main

import (
	"fmt"
	"net/http"

	"log"

	"sync"
)

//
type Middleware func(http.HandlerFunc) http.HandlerFunc

//
func require() Middleware {
	//
	var services = &servicer{mu: sync.RWMutex{}}

	mw_tester(services)

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path

			defer func() {
				if err := recover(); err != nil {
					log.Printf("PANIC: %s: %+v", err, r)
				}
			}()

			if path != "/fix" {
				if err := services.failing(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					log.Printf("error: %s: %+v", err, r)
					fmt.Fprintf(w, "error: %+v: %s", path, err)
					return
				}
			}

			// Questionable whether to all fixing regardless of the state of the database and accounts.

			db := services.database()
			if db == nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				log.Printf("uninitialized database: %+v", r)
				fmt.Fprintln(w, "uninitialized database")
				return
			}

			acc := services.accounts()
			if acc == nil {
				w.WriteHeader(http.StatusServiceUnavailable)
				log.Printf("uninitialized accounts: %+v", r)
				fmt.Fprintln(w, "uninitialized accounts")
				return
			}

			log.Printf("processing: %+v\n", r)
			ctx := &svctx{db: db, acc: acc}
			f(w, r.WithContext(ctx))
			if err := ctx.Err(); err != nil {
				services.initializeFailure(err)
				log.Printf("service error: %s: %+v", err, r)
				fmt.Fprintln(w, err)
				return
			} else {
				services.initializeFailure(nil)
			}
			if err := r.Context().Err(); err != nil {
				log.Printf("request error: %s: %+v", err, r)
				fmt.Fprintln(w, "internal error")
				return
			}
		}
	}
}
