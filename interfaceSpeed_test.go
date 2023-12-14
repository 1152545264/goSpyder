package main

import "testing"

type Sumifier interface{ Add(a, b int32) int32 }

type Sumer struct{ id int32 }

func (match Sumer) Add(a, b int32) int32 { return a + b }

type SumerPointer struct {
	id int32
}

func (match *SumerPointer) Add(a, b int32) int32 {
	return a + b
}

func BenchmarkDirect(b *testing.B) {
	adder := Sumer{id: 6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		adder.Add(10, 12)
	}
}

func BenchMarkInterfacePointer(b *testing.B) {
	adder := &SumerPointer{6754}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Sumifier(adder).Add(10, 12)
	}
}
