package main

import (
	"log"

	"time"
)

//
type s_ struct {
	v, i, x int
}

//
func (s *s_) inc() *int {
	s.v += s.i
	if s.v < 0 {
		return nil
	}
	if s.v > s.x {
		s.v = -s.x
	}
	return &s.v
}

//
func mw_tester(services *servicer) {

	// For setting the states to test responses
	ticker := time.NewTicker(time.Millisecond * 2000)
	go func() {
		db, acc := s_{-10, 3, 20}, s_{-20, 2, 20}
		for range ticker.C {
			log.Printf("+ db:%+v acc:%+v", db, acc)
			services.initializeDatabase(db.inc())
			services.initializeAccounts(acc.inc())
			log.Printf("- db:%+v acc:%+v", db, acc)
		}
	}()
}
