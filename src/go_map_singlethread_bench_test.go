package sharded_map

import "testing"

func BenchmarkGoNativeMap_SingleThread(b *testing.B) {
	benchSinglethreadedGoNativeMap(b)
}
