package sharded_map

import (
	"testing"
)

func BenchmarkShardedMap_Concurrent1(b *testing.B) {
	benchMultithreaded(b, 1)
}
func BenchmarkShardedMap_Concurrent2(b *testing.B) {
	benchMultithreaded(b, 2)
}
func BenchmarkShardedMap_Concurrent4(b *testing.B) {
	benchMultithreaded(b, 4)
}
func BenchmarkShardedMap_Concurrent8(b *testing.B) {
	benchMultithreaded(b, 8)
}
func BenchmarkShardedMap_Concurrent16(b *testing.B) {
	benchMultithreaded(b, 16)
}
func BenchmarkShardedMap_Concurrent32(b *testing.B) {
	benchMultithreaded(b, 32)
}
func BenchmarkShardedMap_Concurrent64(b *testing.B) {
	benchMultithreaded(b, 64)
}
func BenchmarkShardedMap_Concurrent128(b *testing.B) {
	benchMultithreaded(b, 128)
}
func BenchmarkShardedMap_Concurrent256(b *testing.B) {
	benchMultithreaded(b, 256)
}
func BenchmarkShardedMap_Concurrent512(b *testing.B) {
	benchMultithreaded(b, 512)
}
func BenchmarkShardedMap_Concurrent1024(b *testing.B) {
	benchMultithreaded(b, 1024)
}
func BenchmarkShardedMap_Concurrent2048(b *testing.B) {
	benchMultithreaded(b, 2048)
}
func BenchmarkShardedMap_Concurrent4096(b *testing.B) {
	benchMultithreaded(b, 4096)
}
