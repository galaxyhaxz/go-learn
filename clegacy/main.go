// Package clegacy implements functions from the C library to
// make using and porting to Go easier
package clegacy
// package libc

import (
	"unsafe"
)

// MemSet sets an array to the specified byte, not optimized
// Equivalent to memset in C
func MemSet(s unsafe.Pointer, c byte, n uintptr) {
	ptr := uintptr(s)
	var i uintptr
	for i = 0; i < n; i++ {
		pByte := (*byte)(unsafe.Pointer(ptr + i))
		*pByte = c
	}
}

// MemCpy copies a block of memory to another, not optimized
// Equivalent to memcpy in C
func MemCpy(dst unsafe.Pointer, src unsafe.Pointer, n uintptr) {
	d := uintptr(dst)
	s := uintptr(src)
	var i uintptr
	for i = 0; i < n; i++ {
		pDest := (*byte)(unsafe.Pointer(d + i))
		pSrc := (*byte)(unsafe.Pointer(s + i))
		*pDest = *pSrc
	}
}