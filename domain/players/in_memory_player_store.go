package players

import (
	"sync"
)

type PlayerStore interface {
	RecordWin(name string) error
	GetPlayerScore(name string) (int, error)
	GetLeague() ([]Player, error)
}

type Player struct {
	Name string
	Wins int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		map[string]int{},
		sync.RWMutex{},
	}
}

type InMemoryPlayerStore struct {
	store map[string]int
	lock  sync.RWMutex
}

func (i *InMemoryPlayerStore) RecordWin(name string) error {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.store[name]++
	return nil
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	i.lock.RLock()
	defer i.lock.RUnlock()
	return i.store[name], nil
}

func (i *InMemoryPlayerStore) GetLeague() ([]Player, error) {
	var league []Player
	i.lock.RLock()
	for name, wins := range i.store {
		league = append(league, Player{name, wins})
	}
	i.lock.RUnlock()
	return league, nil
}
