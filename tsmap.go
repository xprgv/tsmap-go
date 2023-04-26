// thread safe map

package tsmap

import "sync"

type TsMap[K comparable, V any] struct {
	coreMap map[K]V
	mtx     sync.RWMutex
}

func NewTsMap[K comparable, V any]() TsMap[K, V] {
	return TsMap[K, V]{coreMap: map[K]V{}}
}

func (m *TsMap[K, V]) Get(key K) (value V, exist bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	value, exist = m.coreMap[key]
	return value, exist
}

func (m *TsMap[K, V]) Set(key K, value V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.coreMap[key] = value
}

func (m *TsMap[K, V]) Pop(key K) (value V, exist bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	value, exist = m.coreMap[key]
	if exist {
		delete(m.coreMap, key)
	}
	return value, exist
}

func (m *TsMap[K, V]) Delete(key K) (deleted bool) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	if _, exists := m.coreMap[key]; exists {
		delete(m.coreMap, key)
		return true
	} else {
		return false
	}
}

func (m *TsMap[K, V]) ForEach(f func(value V)) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	for _, value := range m.coreMap {
		f(value)
	}
}

func (m *TsMap[K, V]) Size() int {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	return len(m.coreMap)
}

func (m *TsMap[K, V]) Flush() {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.coreMap = make(map[K]V)
}
