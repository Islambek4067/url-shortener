package storage

import (
	"sync"
	"time"
)

type Memory struct {
	mu    sync.RWMutex
	links map[string]*Link
}

func NewMemory() *Memory {
	return &Memory{
		links: make(map[string]*Link),
	}
}

func (m *Memory) Save(code, url string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.links[code] = &Link{Original: url, CreatedAt: time.Now()}
	return nil
}

func (m *Memory) Get(code string) (Link, bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	l, ok := m.links[code]
	if ok {
		return *l, true, nil
	}
	return Link{}, false, nil
}

func (m *Memory) AddClick(code string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	l, ok := m.links[code]
	if ok {
		l.Clicks++
		return nil
	}
	return nil
}
