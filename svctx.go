package main

import (
	"time"
)

// Local services context provided to handlers
type svctx struct {
	db  *int
	acc *int
	err error
}

// TODO
func (s *svctx) Deadline() (deadline time.Time, ok bool) {
	return
}

// TODO
func (s *svctx) Done() <-chan struct{} {
	return nil
}

//
func (s *svctx) Err() error {
	return s.err
}

//
func (s *svctx) initializeFailure(err error) {
	s.err = err
}

//
func (s *svctx) failing() error {
	return s.err
}

// "db" and "acc" are guaranteed to return a non-nil value since this context will only be initialized after
// the stateful middleware has verified the values are not nil.
func (s *svctx) Value(key interface{}) interface{} {
	switch key {
	case "db", "database":
		return s.db
	case "acc", "accounts":
		return s.acc
	}
	return nil
}
