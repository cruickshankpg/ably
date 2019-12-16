package store

import (
	"sync"
	"time"
)

type SessionState struct {
	Seed           int64
	SequenceLength uint32
}

type Store struct {
	sessions    map[string]SessionState
	sessionLock sync.RWMutex
}

func New() *Store {
	return &Store{
		sessions: make(map[string]SessionState),
	}
}

func (s *Store) Set(key string, val SessionState) {
	s.sessionLock.Lock()
	s.sessions[key] = val
	s.sessionLock.Unlock()
}

func (s *Store) Get(key string) (SessionState, bool) {
	s.sessionLock.RLock()
	val, ok := s.sessions[key]
	s.sessionLock.RUnlock()
	return val, ok
}

func (s *Store) Delete(key string) {
	s.sessionLock.Lock()
	delete(s.sessions, key)
	s.sessionLock.Unlock()
}

func (s *Store) Expire(key string, after time.Duration) {
	//TODO
}
