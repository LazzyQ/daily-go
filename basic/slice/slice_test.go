package slice

import (
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	s := make([]byte, 128)
	t.Logf("len: %d, cap: %d", len(s), cap(s))

	s1 := s[:64]
	t.Logf("len: %d, cap: %d", len(s1), cap(s1))

	s2 := s1[:cap(s1)]
	t.Logf("len: %d, cap: %d", len(s2), cap(s2))

	t.Logf("s: %p, s1: %p, s2: %p", &s, &s1, &s2)
	t.Logf("s: %T, s1: %T, s2: %T", s, s1, s2)

	a := uintptr(unsafe.Pointer(&s1))
	a = a + 8
	t.Logf("%v", *(*int)(unsafe.Pointer(a)))
	a = a + 8
	t.Logf("%v", *(*int)(unsafe.Pointer(a)))

	b := unsafe.Pointer(&s)
	b1 := unsafe.Pointer(&s1)
	b2 := unsafe.Pointer(&s2)

	t.Logf("%v, %v, %v", *(*int)(b), *(*int)(b1), *(*int)(b2))
}
