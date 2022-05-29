package sharded_map

import (
	"reflect"
	"sync"
	"testing"
)

type Key int
type Value int

func keyHasher(k Key) int { return int(k) }

func TestNewShardedMap(t *testing.T) {
	type args struct {
		shardsNum int
		hasher    func(Key) int
	}
	tests := []struct {
		name string
		args args
		want *ShardedMap[Key, Value]
	}{
		{
			name: "-1 shards",
			args: args{
				shardsNum: -1,
				hasher:    keyHasher,
			},
			want: nil,
		},
		{
			name: "0 shards",
			args: args{
				shardsNum: 0,
				hasher:    keyHasher,
			},
			want: nil,
		},
		{
			name: "1 shard",
			args: args{
				shardsNum: 1,
				hasher:    keyHasher,
			},
			want: &ShardedMap[Key, Value]{
				shardsNum: 1,
			},
		},
		{
			name: "999 shards",
			args: args{
				shardsNum: 999,
				hasher:    keyHasher,
			},
			want: &ShardedMap[Key, Value]{
				shardsNum: 999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewShardedMap[Key, Value](tt.args.shardsNum, tt.args.hasher)
			if got == nil && tt.want == nil {
				return
			}
			if got.shardsNum != tt.want.shardsNum {
				t.Errorf("got = %v, want %v", got.shardsNum, tt.want.shardsNum)
			}
			if got.shards == nil {
				t.Errorf("got = %v, want %v", nil, got.shards)
			}
		})
	}
}

func TestShardedMap_Get(t *testing.T) {
	m := NewShardedMap[Key, Value](3, keyHasher)
	m.Set(1, 11)
	m.Set(2, 22)
	m.Set(3, 33)

	type args struct {
		key Key
	}
	tests := []struct {
		name   string
		args   args
		want   Value
		wantOk bool
	}{
		{
			name: "good",
			args: args{
				key: 1,
			},
			want:   11,
			wantOk: true,
		},
		{
			name: "good",
			args: args{
				key: 2,
			},
			want:   22,
			wantOk: true,
		},
		{
			name: "bad",
			args: args{
				key: 4,
			},
			want:   0,
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := m.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}

func TestShardedMap_Set(t *testing.T) {
	m := NewShardedMap[Key, Value](3, keyHasher)

	type args struct {
		key Key
		v   Value
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "insert unique key",
			args: args{
				key: 42,
				v:   123,
			},
		},
		{
			name: "override key",
			args: args{
				key: 42,
				v:   -999,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.Set(tt.args.key, tt.args.v)
		})
	}
}

func TestShardedMap_getHashForKey(t *testing.T) {
	type fields struct {
		shardsNum int
		hasher    func(Key) int
	}
	type args struct {
		key Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "shards1-0",
			fields: fields{
				shardsNum: 1,
				hasher:    keyHasher,
			},
			args: args{
				key: 0,
			},
			want: 0,
		},
		{
			name: "shards1-1",
			fields: fields{
				shardsNum: 1,
				hasher:    keyHasher,
			},
			args: args{
				key: 1,
			},
			want: 0,
		},
		{
			name: "shards1-999",
			fields: fields{
				shardsNum: 1,
				hasher:    keyHasher,
			},
			args: args{
				key: 999,
			},
			want: 0,
		},
		{
			name: "shards2-0",
			fields: fields{
				shardsNum: 2,
				hasher:    keyHasher,
			},
			args: args{
				key: 0,
			},
			want: 0,
		},
		{
			name: "shards2-0",
			fields: fields{
				shardsNum: 2,
				hasher:    keyHasher,
			},
			args: args{
				key: 1,
			},
			want: 1,
		},
		{
			name: "shards2-2",
			fields: fields{
				shardsNum: 2,
				hasher:    keyHasher,
			},
			args: args{
				key: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewShardedMap[Key, Value](tt.fields.shardsNum, tt.fields.hasher)
			if got := m.getHashForKey(tt.args.key); got != tt.want {
				t.Errorf("getHashForKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shard_Get(t *testing.T) {
	type fields struct {
		mutex sync.RWMutex
		data  map[Key]Value
	}
	type args struct {
		key Key
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Value
		wantOk bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &shard[Key, Value]{
				mutex: tt.fields.mutex,
				data:  tt.fields.data,
			}
			got, got1 := s.Get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.wantOk {
				t.Errorf("Get() got1 = %v, want %v", got1, tt.wantOk)
			}
		})
	}
}

func Test_shard_String(t *testing.T) {
	type fields struct {
		data map[Key]Value
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "uninitialized map",
			fields: fields{
				data: nil,
			},
			want: "[]",
		},
		{
			name: "0 elements",
			fields: fields{
				data: map[Key]Value{},
			},
			want: "[]",
		},
		{
			name: "a few elements",
			fields: fields{
				data: map[Key]Value{1: 11, 2: 22},
			},
			want: "[1:11, 2:22, ]",
		},
		{
			name: "a lot of elements",
			fields: fields{
				data: map[Key]Value{0: 0, 1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6, 7: 7, 8: 8, 9: 9, 10: 10, 11: 11, 12: 12, 13: 13, 14: 14, 15: 15, 16: 16, 17: 17, 18: 18, 19: 19, 20: 20, 21: 21},
			},
			want: "[ ...22 items... ]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := shard[Key, Value]{
				data: tt.fields.data,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
