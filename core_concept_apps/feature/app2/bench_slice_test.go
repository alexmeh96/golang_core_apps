package main

import "testing"

func nonAllocSlic(n int) {
	ints := []int{}
	for i := 0; i < n; i++ {
		ints = append(ints, i)
	}
}

func allocSlice(n int) {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = i
	}
}

func BenchmarkSliceNonAlloc(b *testing.B) {
	n := 100

	for i := 0; i < b.N; i++ {
		nonAllocSlic(n)
	}
}

func BenchmarkSliceAlloc(b *testing.B) {
	n := 100

	for i := 0; i < b.N; i++ {
		allocSlice(n)
	}
}

func BenchmarkSlice(b *testing.B) {
	n := 100

	b.Run("non alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			nonAllocSlic(n)
		}
	})

	b.Run("alloc", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			allocSlice(n)
		}
	})
}
