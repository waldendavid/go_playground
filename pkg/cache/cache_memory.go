package cache

import "sync"

type Key string

type Value interface{}

type MemoryCache struct {
	data     map[Key]Value
	threaded bool
	lock     sync.RWMutex
}

func NewMemoryCache(options ...func(*MemoryCache)) *MemoryCache {
	mc := &MemoryCache{}
	for _, o := range options {
		o(mc)
	}
	return &MemoryCache{
		data: make(map[Key]Value),
	}
}

func WithThreaded(t bool) func(*MemoryCache) {
	return func(mc *MemoryCache) {
		mc.threaded = t
	}
}

func (mc *MemoryCache) Get(k Key) (Value, bool) {
	if mc.threaded {
		//todo
		mc.lock.RLock()
		defer mc.lock.RUnlock()
	}
	value, exists := mc.data[k]
	if !exists {
		return nil, false
	}
	return value, true
}

func (mc *MemoryCache) Set(k Key, v Value) {
	if mc.threaded {
		//todo
		mc.lock.Lock()
		defer mc.lock.Unlock()
	}
	mc.data[k] = v
}

func (mc *MemoryCache) Remove(k Key) {
	if mc.threaded {
		mc.lock.Lock()
		defer mc.lock.Unlock()
	}
	delete(mc.data, k)
}

// najpierw czy jest w cache a potem dopiero zapytanie
