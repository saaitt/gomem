package domain

import (
	"fmt"
	"sync"
	"time"
)

var NotFoundError = fmt.Errorf("not_found")

type TTLRadixNode struct {
	children  map[byte]*TTLRadixNode
	value     *string
	ttl       *time.Duration
	expiredAt time.Time
	isEnd     bool
}

type TTLRadixTree struct {
	root *TTLRadixNode
	mu   sync.RWMutex
}

func (rn *TTLRadixNode) InsertNode(key *byte) {
	if rn.children == nil {
		rn.children = make(map[byte]*TTLRadixNode)
	}
	rn.children[*key] = &TTLRadixNode{}
}

func (rn *TTLRadixNode) MakeEdge(value *string, ttl *time.Duration) {
	rn.value = value
	rn.ttl = ttl
	rn.expiredAt = time.Now().UTC().Add(*ttl)
	//rn.isEnd = true
}

func (rt *TTLRadixTree) Insert(key, value *string, ttl *time.Duration) {
	rt.mu.Lock()
	keyBytes := []byte(*key)
	current := rt.root
	if current == nil {
		rt.root = &TTLRadixNode{}
		current = rt.root
	}
	for _, kbItem := range keyBytes {
		if current.children != nil {
			child, ok := current.children[kbItem]
			if ok {
				current = child
				continue
			}
		}
		current.InsertNode(&kbItem)
		current = current.children[kbItem]
	}
	current.MakeEdge(value, ttl)
	rt.mu.Unlock()
}

func (rt *TTLRadixTree) Find(key string) (string, error) {
	rt.mu.RLock()
	defer rt.mu.RUnlock()

	keyBytes := []byte(key)
	current := rt.root
	if current == nil {
		rt.root = &TTLRadixNode{}
		current = rt.root
	}
	for _, kbItem := range keyBytes {
		if current.children != nil {
			_, ok := current.children[kbItem]
			if !ok {
				return "", NotFoundError
			}
			current = current.children[kbItem]
			continue
		}
		return "", NotFoundError
	}
	return *current.value, nil
}
