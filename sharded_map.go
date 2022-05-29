package sharded_map

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
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
	for k, v := range s.data {
		sb.WriteString(fmt.Sprintf("%v:%v, ", k, v))
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
	shardsNum int
	hasher    func(K) int
	shards    []*shard[K, V]
}

func NewShardedMap[K comparable, V any](shardsNum int, hasher func(K) int) *ShardedMap[K, V] {
	if shardsNum < 1 {
		return nil
	}

	m := &ShardedMap[K, V]{
		shardsNum: shardsNum,
		hasher:    hasher,
		shards:    make([]*shard[K, V], shardsNum),
	}

	for i := 0; i < shardsNum; i++ {
		m.shards[i] = newShard[K, V]()
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
	s := m.shards[hash]
	return s
}

func (m *ShardedMap[K, V]) getHashForKey(k K) int {
	hash := m.hasher(k)
	return hash % m.shardsNum
}

var _ Map[int, int] = (*ShardedMap[int, int])(nil)
