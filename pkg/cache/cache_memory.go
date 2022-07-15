package cache

import "sync"

type Key string

type Value interface{}

type MemoryCache struct {
	data map[Key]Value
	lock sync.RWMutex
}

func NewMemoryCache() *MemoryCache {
	// lock - autoinicjalizacja
	return &MemoryCache{
		data: make(map[Key]Value),
	}
}

const Threaded = true

func (mc *MemoryCache) Get(k Key) (Value, bool) {
	if Threaded {
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
	if Threaded {
		//todo
		mc.lock.Lock()
		defer mc.lock.Unlock()
	}
	mc.data[k] = v
}

func (mc *MemoryCache) Remove(k Key) {
	if Threaded {
		mc.lock.Lock()
		defer mc.lock.Unlock()
	}
	delete(mc.data, k)
}

// najpierw czy jest w cache a potem dopiero zapytanie
