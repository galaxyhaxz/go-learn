// Package clegacy implements functions from the C library to
// make using and porting to Go easier
package clegacy

import (
	"unsafe"
)

// Set an array to the specified byte, not optimized
func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}
