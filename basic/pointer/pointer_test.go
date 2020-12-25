package pointer

import (
	"sync"
	"testing"
	"unsafe"
)

func TestBasePointer(t *testing.T) {
	num := 1
	numPointer := &num
	t.Logf("num的值: %v, num的指针: %v, num的指针: %p, 通过指针获取值: %v", num, numPointer, numPointer, *numPointer)
}

func TestUintptr(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 0, 1, 2)
	t.Logf("len: %d, cap: %d", len(s), cap(s))

	sp := unsafe.Pointer(&s)
	sa := uintptr(sp)
	sa = sa + unsafe.Sizeof(sp)
	t.Logf("len: %d", *(*int)(unsafe.Pointer(sa)))
	sa = sa + unsafe.Sizeof(int(0))
	t.Logf("cap: %d", *(*int)(unsafe.Pointer(sa)))
}

func TestUnsafePointerBridge(t *testing.T) {
	var vInt64 int64 = 1
	p := unsafe.Pointer(&vInt64)
	var i int = *(*int)(p)
	t.Logf("%v", i)
}

func BenchmarkSyncPoolSlice(b *testing.B) {
	p := sync.Pool{
		New: func() interface{} {
			return make([]byte, 128)
		},
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bytes, _ := p.Get().([]byte)
			bytes[0] = 1
			p.Put(bytes)
		}
	})

}

func BenchmarkSyncPoolSlicePointer(b *testing.B) {
	p := sync.Pool{
		New: func() interface{} {
			bytes := make([]byte, 128)
			return &bytes
		},
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bytes, _ := p.Get().(*[]byte)
			(*bytes)[0] = 1
			p.Put(bytes)
		}
	})

}
