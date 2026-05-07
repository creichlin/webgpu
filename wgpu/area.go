package wgpu

import (
	"sync/atomic"
)

type allocChunk[T any] struct {
	Next atomic.Uint32

	// avoid false sharing by moving the actual data
	// to the next cache line
	_ [60]byte

	Values [1024]T
}

type allocArea[T any] struct {
	chunk atomic.Pointer[allocChunk[T]]
}

func (a *allocArea[T]) Get() *T {
	chunk := a.chunk.Load()

	if chunk == nil {
		// there was no chunk,
		// initialize first chunk and set active
		chunk = new(allocChunk[T])
		a.chunk.Store(chunk)
	}

	// get the next slot in the active chunk
	// and claim the active one directly
	next := chunk.Next.Add(1) - 1

	if int(next) >= len(chunk.Values) {
		// chunk is oom,
		// initialize new chunk and set active
		chunk = new(allocChunk[T])
		chunk.Next.Store(1)
		a.chunk.Store(chunk)
		next = 0
	}

	// take the address of the value
	return &chunk.Values[next]
}
