package gorillaws

import (
	"sync"

	port "github.com/muhriddinnorqulov/skeleton/src/core/domain/ports/wsport"
)

type ConnectionStore struct {
	mu    sync.RWMutex
	items map[port.ConnectionID]*Connection
}

// @inject
func NewConnectionStore() *ConnectionStore {
	return &ConnectionStore{items: make(map[port.ConnectionID]*Connection)}
}

func (this *ConnectionStore) Add(c *Connection) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.items[c.id] = c
}

func (this *ConnectionStore) Get(id port.ConnectionID) (*Connection, bool) {
	this.mu.RLock()
	defer this.mu.RUnlock()
	c, ok := this.items[id]
	return c, ok
}

func (this *ConnectionStore) Remove(id port.ConnectionID) {
	this.mu.Lock()
	defer this.mu.Unlock()
	delete(this.items, id)
}

func (this *ConnectionStore) All() []*Connection {
	this.mu.RLock()
	defer this.mu.RUnlock()
	list := make([]*Connection, 0, len(this.items))
	for _, c := range this.items {
		list = append(list, c)
	}
	return list
}
