package players

import "sync"

type PlayerStore interface {
	RecordWin(name string) error
	GetPlayerScore(name string) (int, error)
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
