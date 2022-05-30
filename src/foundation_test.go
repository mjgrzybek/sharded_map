package sharded_map

import (
	"runtime"
	"sync"
	"testing"
)

func benchMultithreaded(b *testing.B, shards int) {
	threads := runtime.GOMAXPROCS(0)
	run(b, threads, func() Map[int, int] {
		return NewShardedMap[int, int](shards, intHasher)
	})
}

func benchSinglethreadedShardedMap(b *testing.B, shards int) {
	const threads = 1
	run(b, threads, func() Map[int, int] {
		return NewShardedMap[int, int](shards, intHasher)
	})
}

func benchSinglethreadedGoNativeMap(b *testing.B) {
	const threads = 1
	run(b, threads, func() Map[int, int] {
		return NewGoMap[int, int]()
	})
}

func run(b *testing.B, threads int, newMap func() Map[int, int]) {
	for i := 0; i < b.N; i++ {
		m := newMap()
		floodWithData(m, 1024*1024, threads)
	}
}

func floodWithData(m Map[int, int], numberOfEntriesPerRoutine int, routinesNumber int) {
	wg := sync.WaitGroup{}
	wg.Add(routinesNumber)
	for i := 0; i < routinesNumber; i++ {
		go func(i int, wg *sync.WaitGroup) {
			for j := i * numberOfEntriesPerRoutine; j < (i+1)*numberOfEntriesPerRoutine; j++ {
				v := i * j
				m.Set(v, v)
			}
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

func intHasher(v int) int { return v }

type GoMap[K comparable, V any] struct {
	m map[K]V
}

func NewGoMap[K comparable, V any]() *GoMap[K, V] {
	return &GoMap[K, V]{m: make(map[K]V)}
}

func (m *GoMap[K, V]) Set(k K, v V) {
	m.m[k] = v
}

func (m *GoMap[K, V]) Get(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}
