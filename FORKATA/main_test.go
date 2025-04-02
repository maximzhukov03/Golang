package main

import (
	"fmt"
	"testing"
)

func BenchmarkCrc64(b *testing.B){
	bench := NewHashMap(16, WithHashCRC64())
	for i := 0; i < b.N; i++{
		bench.Set(fmt.Sprintf("key%d", i), i)
	}
}

func BenchmarkCrc32(b *testing.B){
	bench := NewHashMap(16, WithHashCRC32())
	for i := 0; i < b.N; i++{
		bench.Set(fmt.Sprintf("key%d", i), i)
	}
}

func BenchmarkCrc16(b *testing.B){
	bench := NewHashMap(16, WithHashCRC16())
	for i := 0; i < b.N; i++{
		bench.Set(fmt.Sprintf("key%d", i), i)
	}
}

func BenchmarkCrc8(b *testing.B){
	bench := NewHashMap(16, WithHashCRC8())
	for i := 0; i < b.N; i++{
		bench.Set(fmt.Sprintf("key%d", i), i)
	}
}