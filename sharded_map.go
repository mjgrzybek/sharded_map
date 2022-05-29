package sharded_map

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/exp/maps"
)

type Map[K comparable, V any] interface {
	Set(K, V)
	Get(K) (V, bool)
}

type shard[K comparable, V any] struct {
	mutex sync.RWMutex
	data  map[K]V
}

func (s shard[K, V]) String() string {
	dataLen := len(s.data)
	if dataLen > 20 {
		return "[ ..." + strconv.Itoa(dataLen) + " items... ]"
	}

	sb := strings.Builder{}
	vals := maps.Values(s.data)
	for _, v := range vals {
		sb.WriteString(fmt.Sprintf("%v, ", v))
	}
	return "[" + sb.String() + "]"
}

func (s *shard[K, V]) Get(k K) (V, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	v, ok := s.data[k]
	return v, ok
}

func (s *shard[K, V]) Set(k K, v V) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.data[k] = v
}

func newShard[K comparable, V any]() *shard[K, V] {
	return &shard[K, V]{
		data: make(map[K]V),
	}
}

type ShardedMap[K comparable, V any] struct {
	shardsLimit int
	hasher      func(K) int
	shardsMap   map[int]*shard[K, V]
}

func NewShardedMap[K comparable, V any](shardsLimit int, hasher func(K) int) *ShardedMap[K, V] {
	m := &ShardedMap[K, V]{
		shardsLimit: shardsLimit,
		hasher:      hasher,
		shardsMap:   make(map[int]*shard[K, V]),
	}

	for i := 0; i < shardsLimit; i++ {
		m.createShardIfNotExist(i)
	}

	return m
}

func (m *ShardedMap[K, V]) Set(k K, v V) {
	shard := m.getShardForKey(k)
	shard.Set(k, v)
}

func (m *ShardedMap[K, V]) Get(k K) (V, bool) {
	shard := m.getShardForKey(k)

	v, ok := shard.Get(k)
	return v, ok
}

func (m *ShardedMap[K, V]) getShardForKey(k K) *shard[K, V] {
	hash := m.getHashForKey(k)
	s := m.shardsMap[hash]
	return s
}

func (m *ShardedMap[K, V]) createShardIfNotExist(hash int) {
	_, ok := m.shardsMap[hash]
	if !ok {
		m.shardsMap[hash] = newShard[K, V]()
	}
}

func (m *ShardedMap[K, V]) getHashForKey(k K) int {
	hash := m.hasher(k)
	return hash % m.shardsLimit
}

var _ Map[int, int] = (*ShardedMap[int, int])(nil)
