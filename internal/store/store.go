package store

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

type SessionState struct {
	Seed           int64
	SequenceLength uint32
}

type Store struct {
	//sessions    map[string]SessionState
	sessions *cache.Cache
}

func New() *Store {
	return &Store{
		sessions: cache.New(time.Second*30, time.Second*5),
	}
}

func (s *Store) Set(key string, val SessionState) {
	s.sessions.Set(key, val, cache.DefaultExpiration)
}

func (s *Store) Get(key string) (SessionState, bool) {
	o, ok := s.sessions.Get(key)
	if !ok {
		return SessionState{}, ok
	}

	val := o.(SessionState)

	return val, ok
}

func (s *Store) Delete(key string) {
	s.sessions.Delete(key)
}
