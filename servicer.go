package main

import (
	"sync"
)

// Global state uf services required by endpoints
type servicer struct {
	db  *int
	acc *int
	err error
	mu  sync.RWMutex
}

//
func (s *servicer) initializeDatabase(db *int) *int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.db = db
	return s.db
}

//
func (s *servicer) initializeAccounts(acc *int) *int {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.acc = acc
	return s.acc
}

//
func (s *servicer) initializeFailure(err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.err = err
}

//
func (s *servicer) failing() error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.err
}

//
func (s *servicer) database() *int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.db
}

//
func (s *servicer) accounts() *int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.acc
}
