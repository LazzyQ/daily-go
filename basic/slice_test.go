package basic

import (
	"testing"
	"unsafe"
)

func TestSlice(t *testing.T) {
	// 如果初始化时只有一个参数，那么len，cap都是这个值
	s := make([]byte, 128)
	t.Logf("len: %d, cap: %d", len(s), cap(s)) //  len: 128, cap: 128

	// 对slice进行切片，
	s1 := s[:64]
	t.Logf("len: %d, cap: %d", len(s1), cap(s1)) // len: 64, cap: 128

	s2 := s1[:cap(s1)]
	t.Logf("len: %d, cap: %d", len(s2), cap(s2)) // len: 128, cap: 128

	s3 := s[10:20]
	t.Logf("len: %d, cap: %d", len(s3), cap(s3)) // len: 10, cap: 118

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

func TestSliceStructChange(t *testing.T) {
	students := make([]Student, 0, 2)

	student1 := Student{Name: "x", Age: 1}
	student2 := Student{Name: "y", Age: 2}
	students = append(students, student1, student2)

	// changeStudents(students)
	// t.Log(students) // [{x 100} {y 200}]

	copyStudents := make([]Student, 0, len(students))
	copyStudents = append(copyStudents, students...)

	changeStudents(copyStudents)
	t.Log(students)
	t.Log(copyStudents)
}

func changeStudents(students []Student) {
	students[0].Age = 100
	students[1].Age = 200
}

func TestSliceStructChange2(t *testing.T) {
	students := make([]Student, 0, 2)

	student1 := Student{Name: "x", Age: 1}
	student2 := Student{Name: "y", Age: 2}
	students = append(students, student1, student2)

	t.Log(students)

	// change student1
	student1.Name = "xx"
	t.Log(students)

	// change student1
	students[0].Name = "xx"
	t.Log(students)
}

func TestSliceForChangeValue(t *testing.T) {

	students := make([]Student, 0, 2)

	student1 := Student{Name: "x", Age: 1, Parents: map[string]string{"mom": "momName1", "dady": "dadyName1"}}
	student2 := Student{Name: "y", Age: 2, Parents: map[string]string{"mom": "momName2", "dady": "dadyName2"}}
	students = append(students, student1, student2)

	for index, student := range students {
		t.Logf("%p\n", &student)

		student.Name = student.Name + "changed"
		student.Parents["mom"] = student.Parents["mom"] + "changed"
		students[index] = student
	}

	for _, student := range students {
		t.Logf("%+v", student)
	}
}

func TestSlicePointPrint(t *testing.T) {
	student1 := Student{
		Name: "zhangsan",
		Age:  15,
	}

	student2 := Student{
		Name: "lisi",
		Age:  16,
	}

	studentStructs := []Student{student1, student2}
	studentPointers := []*Student{&student1, &student2}

	t.Logf("studentStructs: %v", studentStructs)
	t.Logf("studentStructs: %+v", studentStructs)

	t.Logf("studentPointers: %v", studentPointers)
	t.Logf("studentPointers: %+v", studentPointers)
}
