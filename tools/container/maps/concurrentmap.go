// Copyright 2021 dudaodong@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license

// Package maps includes some functions to manipulate map.
package maps

import (
	"fmt"
	"sync"
	"unsafe"
)

const defaultShardCount = 32

// ConcurrentMap is like map, but is safe for concurrent use by multiple goroutines.
type ConcurrentMap[K comparable, V any] struct {
	shardCount uint64
	locks      []sync.RWMutex
	maps       []map[K]V
}

// NewConcurrentMap create a ConcurrentMap with specific shard count.
// Play: https://go.dev/play/p/3PenTPETJT0
func NewConcurrentMap[K comparable, V any](shardCount int) *ConcurrentMap[K, V] {
	if shardCount <= 0 {
		shardCount = defaultShardCount
	}

	cm := &ConcurrentMap[K, V]{
		shardCount: uint64(shardCount),
		locks:      make([]sync.RWMutex, shardCount),
		maps:       make([]map[K]V, shardCount),
	}

	for i := range cm.maps {
		cm.maps[i] = make(map[K]V)
	}

	return cm
}

// Set the value for a key.
// Play: https://go.dev/play/p/3PenTPETJT0
func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	shard := cm.getShard(key)

	cm.locks[shard].Lock()
	cm.maps[shard][key] = value

	cm.locks[shard].Unlock()
}

// Get the value stored in the map for a key, or nil if no.
// Play: https://go.dev/play/p/3PenTPETJT0
func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	value, ok := cm.maps[shard][key]
	cm.locks[shard].RUnlock()

	return value, ok
}

// GetOrSet returns the existing value for the key if present.
// Otherwise, it sets and returns the given value.
// Play: https://go.dev/play/p/aDcDApOK01a
func (cm *ConcurrentMap[K, V]) GetOrSet(key K, value V) (actual V, ok bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	if actual, ok := cm.maps[shard][key]; ok {
		cm.locks[shard].RUnlock()
		return actual, ok
	}
	cm.locks[shard].RUnlock()

	// lock again
	cm.locks[shard].Lock()
	if actual, ok = cm.maps[shard][key]; ok {
		cm.locks[shard].Unlock()
		return
	}

	cm.maps[shard][key] = value
	cm.locks[shard].Unlock()

	return value, ok
}

// Delete the value for a key.
// Play: https://go.dev/play/p/uTIJZYhpVMS
func (cm *ConcurrentMap[K, V]) Delete(key K) {
	shard := cm.getShard(key)

	cm.locks[shard].Lock()
	delete(cm.maps[shard], key)
	cm.locks[shard].Unlock()
}

// GetAndDelete returns the existing value for the key if present and then delete the value for the key.
// Otherwise, do nothing, just return false
// Play: https://go.dev/play/p/ZyxeIXSZUiM
func (cm *ConcurrentMap[K, V]) GetAndDelete(key K) (actual V, ok bool) {
	shard := cm.getShard(key)

	cm.locks[shard].RLock()
	if actual, ok = cm.maps[shard][key]; ok {
		cm.locks[shard].RUnlock()
		cm.Delete(key)
		return
	}
	cm.locks[shard].RUnlock()

	return actual, false
}

// Has checks if map has the value for a key.
// Play: https://go.dev/play/p/C8L4ul9TVwf
func (cm *ConcurrentMap[K, V]) Has(key K) bool {
	_, ok := cm.Get(key)
	return ok
}

// Range calls iterator sequentially for each key and value present in each of the shards in the map.
// If iterator returns false, range stops the iteration.
// Play: https://go.dev/play/p/iqcy7P8P0Pr
func (cm *ConcurrentMap[K, V]) Range(iterator func(key K, value V) bool) {
	for shard := range cm.locks {
		cm.locks[shard].RLock()

		for k, v := range cm.maps[shard] {
			if !iterator(k, v) {
				cm.locks[shard].RUnlock()
				return
			}
		}
		cm.locks[shard].RUnlock()
	}
}

// getShard get shard by a key.
func (cm *ConcurrentMap[K, V]) getShard(key K) uint64 {
	hash := hashKey(key)
	return hash % cm.shardCount
}

// hashKey computes a hash for the given key, with fast paths for common types
// to avoid the overhead of fmt.Sprintf.
func hashKey[K comparable](key K) uint64 {
	switch k := any(key).(type) {
	case string:
		return fnv64String(k)
	case int:
		return fnv64Uint64(uint64(k))
	case int8:
		return fnv64Uint64(uint64(k))
	case int16:
		return fnv64Uint64(uint64(k))
	case int32:
		return fnv64Uint64(uint64(k))
	case int64:
		return fnv64Uint64(uint64(k))
	case uint:
		return fnv64Uint64(uint64(k))
	case uint8:
		return fnv64Uint64(uint64(k))
	case uint16:
		return fnv64Uint64(uint64(k))
	case uint32:
		return fnv64Uint64(uint64(k))
	case uint64:
		return fnv64Uint64(k)
	case uintptr:
		return fnv64Uint64(uint64(k))
	default:
		return fnv64String(fmt.Sprintf("%v", key))
	}
}

const (
	fnvOffset64 = uint64(14695981039346656037)
	fnvPrime64  = uint64(1099511628211)
)

func fnv64String(key string) uint64 {
	hash := fnvOffset64
	ptr := unsafe.StringData(key)
	for i := 0; i < len(key); i++ {
		hash ^= uint64(*ptr)
		hash *= fnvPrime64
		ptr = (*byte)(unsafe.Add(unsafe.Pointer(ptr), 1))
	}
	return hash
}

func fnv64Uint64(val uint64) uint64 {
	b := [8]byte{
		byte(val), byte(val >> 8), byte(val >> 16), byte(val >> 24),
		byte(val >> 32), byte(val >> 40), byte(val >> 48), byte(val >> 56),
	}
	hash := fnvOffset64
	for _, c := range b {
		hash ^= uint64(c)
		hash *= fnvPrime64
	}
	return hash
}
