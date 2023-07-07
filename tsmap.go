// thread safe map

package tsmap

import "sync"

type ThreadSafeMap[K comparable, V any] struct {
	coreMap map[K]V
	mu      sync.RWMutex
}

func NewThreadSafeMap[K comparable, V any]() ThreadSafeMap[K, V] {
	return ThreadSafeMap[K, V]{coreMap: map[K]V{}}
}

func (m *ThreadSafeMap[K, V]) Get(key K) (value V, exist bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exist = m.coreMap[key]
	return value, exist
}

func (m *ThreadSafeMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.coreMap[key] = value
}

func (m *ThreadSafeMap[K, V]) Pop(key K) (value V, exist bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, exist = m.coreMap[key]
	if exist {
		delete(m.coreMap, key)
	}
	return value, exist
}

func (m *ThreadSafeMap[K, V]) Delete(key K) (deleted bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.coreMap[key]; exists {
		delete(m.coreMap, key)
		return true
	} else {
		return false
	}
}

func (m *ThreadSafeMap[K, V]) DeleteMultiple(keys ...K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, key := range keys {
		delete(m.coreMap, key)
	}
}

func (m *ThreadSafeMap[K, V]) ForEach(f func(value V)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for _, value := range m.coreMap {
		f(value)
	}
}

func (m *ThreadSafeMap[K, V]) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.coreMap)
}

func (m *ThreadSafeMap[K, V]) Flush() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.coreMap = make(map[K]V)
}
