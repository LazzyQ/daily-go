package basic

import (
	"sort"
	"strings"
	"testing"
	"unsafe"
)

var (
	byteData = []byte("一去二三里，烟村四五家")
)

func BenchmarkSlice2String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = string(byteData)
	}

}

func BenchmarkSlice2String2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SliceByteToString(byteData)
	}
}

func SliceByteToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func TestStringBuilder(t *testing.T) {
	b := strings.Builder{}
	b.Write([]byte{97, 97})
	t.Log(b.String())
}

func TestStringSort(t *testing.T) {
	strs := make([]string, 0, 10)
	strs = append(strs, "1", "2", "3", "4")

	sort.Strings(strs)

	for _, s := range strs {
		t.Log(s)
	}

	t.Log(strs)
	s := make([]string, 0, 20)
	s = append(s, strs...)
	t.Log(len(s))
}
