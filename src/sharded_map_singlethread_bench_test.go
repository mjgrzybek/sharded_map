package sharded_map

import "testing"

func BenchmarkShardedMap_SingleThread1(b *testing.B) {
	benchSinglethreadedShardedMap(b, 1)
}
func BenchmarkShardedMap_SingleThread2(b *testing.B) {
	benchSinglethreadedShardedMap(b, 2)
}
func BenchmarkShardedMap_SingleThread4(b *testing.B) {
	benchSinglethreadedShardedMap(b, 4)
}
func BenchmarkShardedMap_SingleThread8(b *testing.B) {
	benchSinglethreadedShardedMap(b, 8)
}
func BenchmarkShardedMap_SingleThread16(b *testing.B) {
	benchSinglethreadedShardedMap(b, 16)
}
func BenchmarkShardedMap_SingleThread32(b *testing.B) {
	benchSinglethreadedShardedMap(b, 32)
}
func BenchmarkShardedMap_SingleThread64(b *testing.B) {
	benchSinglethreadedShardedMap(b, 64)
}
func BenchmarkShardedMap_SingleThread128(b *testing.B) {
	benchSinglethreadedShardedMap(b, 128)
}
func BenchmarkShardedMap_SingleThread256(b *testing.B) {
	benchSinglethreadedShardedMap(b, 256)
}
func BenchmarkShardedMap_SingleThread512(b *testing.B) {
	benchSinglethreadedShardedMap(b, 512)
}
func BenchmarkShardedMap_SingleThread1024(b *testing.B) {
	benchSinglethreadedShardedMap(b, 1024)
}
func BenchmarkShardedMap_SingleThread2048(b *testing.B) {
	benchSinglethreadedShardedMap(b, 2048)
}
func BenchmarkShardedMap_SingleThread4096(b *testing.B) {
	benchSinglethreadedShardedMap(b, 4096)
}
