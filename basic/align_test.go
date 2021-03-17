package basic

import (
	"testing"
	"unsafe"
)

type T1 struct {
	//24
	A uint8  //1
	B uint16 //2
	E uint8  //1
	C uint32 //4
	D uint64 //8
}
type T2 struct {
	// 16
	A uint8  //1
	E uint8  //1
	B uint16 //2
	C uint32 //4
	D uint64 //8
}

type T3 struct {
	A uint8  //1
	E uint8  //1
	B uint16 //2
}

func Test(t *testing.T) {

	t1 := T1{}
	t2 := T2{}
	t3 := T3{}

	t.Logf("T1 Alignof: %d, SizeOf: %d, ", unsafe.Alignof(t1), unsafe.Sizeof(t1))
	t.Logf("T1.A Alignof: %d, Offsetof: %d, SizeOf: %d, ", unsafe.Alignof(t1.A), unsafe.Offsetof(t1.A), unsafe.Sizeof(t1.A))
	t.Logf("T1.B Alignof: %d, Offsetof: %d, SizeOf: %d, ", unsafe.Alignof(t1.B), unsafe.Offsetof(t1.B), unsafe.Sizeof(t1.B))
	t.Logf("T1.E Alignof: %d, Offsetof: %d, SizeOf: %d, ", unsafe.Alignof(t1.E), unsafe.Offsetof(t1.E), unsafe.Sizeof(t1.E))
	t.Logf("T1.C Alignof: %d, Offsetof: %d, SizeOf: %d, ", unsafe.Alignof(t1.C), unsafe.Offsetof(t1.C), unsafe.Sizeof(t1.C))
	t.Logf("T1.D Alignof: %d, Offsetof: %d, SizeOf: %d, ", unsafe.Alignof(t1.D), unsafe.Offsetof(t1.D), unsafe.Sizeof(t1.D))

	t.Logf("T2 Alignof: %d, SizeOf: %d, ", unsafe.Alignof(t2), unsafe.Sizeof(t2))
	t.Logf("T3 Alignof: %d, SizeOf: %d, ", unsafe.Alignof(t3), unsafe.Sizeof(t3))
}
