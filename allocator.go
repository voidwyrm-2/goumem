package goumem

import "github.com/voidwyrm-2/goumem/allocator"

var mem allocator.MemoryAllocator

func SetMemoryAllocator(m allocator.MemoryAllocator) {
	mem = m
}
