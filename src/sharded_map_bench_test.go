package sharded_map

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func floodWithData(m *ShardedMap[int, int], numberOfEntriesPerRoutine int, routinesNumber int) {
	fmt.Println("Total number of items:", numberOfEntriesPerRoutine*routinesNumber)

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
func BenchmarkShardedMap_128(b *testing.B) {
	impl(b, 128)
}
func BenchmarkShardedMap_256(b *testing.B) {
	impl(b, 256)
}
func BenchmarkShardedMap_512(b *testing.B) {
	impl(b, 512)
}
func BenchmarkShardedMap_1024(b *testing.B) {
	impl(b, 1024)
}
func BenchmarkShardedMap_2048(b *testing.B) {
	impl(b, 2048)
}
func BenchmarkShardedMap_4096(b *testing.B) {
	impl(b, 4096)
}

func impl(b *testing.B, shards int) {
	threads := runtime.GOMAXPROCS(0)
	for i := 0; i < b.N; i++ {
		m := NewShardedMap[int, int](shards, intHasher)
		floodWithData(m, 1024*32*32, threads)
	}
}

func intHasher(v int) int { return v }
