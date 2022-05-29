package sharded_map

import (
	"sync"
	"testing"
	"time"
)

func intHasher(v int) int { return v }

func floodWithData(m *ShardedMap[int, int], numberOfEntries int) {
	floodStart := time.Now().Add(1 * time.Second)

	wg := sync.WaitGroup{}
	for i := 0; i < numberOfEntries; i++ {
		wg.Add(1)
		go func(v int, wg *sync.WaitGroup) {
			sleepDuration := floodStart.Sub(time.Now())
			time.Sleep(sleepDuration)
			m.Set(v, v)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}

func impl(b *testing.B, shards int) {
	for i := 0; i < b.N; i++ {
		m := NewShardedMap[int, int](shards, intHasher)
		floodWithData(m, 1024*64*64+2)
	}
}

func BenchmarkShardedMap_1(b *testing.B) {
	impl(b, 1)
}
func BenchmarkShardedMap_2(b *testing.B) {
	impl(b, 2)
}
func BenchmarkShardedMap_4(b *testing.B) {
	impl(b, 4)
}
func BenchmarkShardedMap_8(b *testing.B) {
	impl(b, 8)
}
func BenchmarkShardedMap_16(b *testing.B) {
	impl(b, 16)
}
func BenchmarkShardedMap_32(b *testing.B) {
	impl(b, 32)
}
func BenchmarkShardedMap_64(b *testing.B) {
	impl(b, 64)
}
