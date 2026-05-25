//go:build wasip2

package main

import (
	"sync"

	"go.bytecodealliance.org/cm"
)

type resourceMap[T any] struct {
	map_  map[cm.Rep]T
	index cm.Rep
	mu    sync.Mutex
}

func (r *resourceMap[T]) Add(selfGo T) cm.Rep {
	r.mu.Lock()
	defer r.mu.Unlock()
	i := r.index
	r.index++
	r.map_[i] = selfGo
	return i
}

func (r *resourceMap[T]) Get(self cm.Rep) T {
	return r.map_[self]
}

func main() {}
