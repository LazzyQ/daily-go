package basic

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type Compose struct {
	Num int
}

func TestMapKey(t *testing.T) {
	a := &Compose{Num: 1}
	b := &Compose{Num: 1}

	m := make(map[*Compose]string)

	m[a] = "a"
	m[b] = "b"

	t.Log(len(m))
}

func TestConcurrentMapReadWrite(t *testing.T) {
	m := make(map[string]string)

	m["s"] = "s"
	for {
		if v, ok := m["s"]; ok {
			fmt.Println(v)
		} else {
			fmt.Println("==")
		}
	}
}

func MapWrite(m *sync.Map) {
	for {
		m.Store("s", "s")
		time.Sleep(time.Millisecond)
	}
}

func TestMapLen(t *testing.T) {
	m := make(map[string]struct{}, 10)
	t.Logf("len(m): %d", len(m))

	m["a"] = struct{}{}
	m["b"] = struct{}{}
	t.Logf("len(m): %d", len(m))
}

func TestMapKey2(t *testing.T) {
	m := make(map[string]string)
	m["A"] = "a"
	m["B"] = "b"

	for k := range m {
		m[k] = k
	}

	for k := range m {
		t.Log(k)
	}
}

func TestMapGrow(t *testing.T) {
	outer := make(map[string]map[string]struct{})
	outer["a"] = make(map[string]struct{})

	inner := outer["a"]
	t.Log(len(inner))
	inner["a"] = struct{}{}
	inner["b"] = struct{}{}
	inner["c"] = struct{}{}
	inner["d"] = struct{}{}
	inner["e"] = struct{}{}
	inner["f"] = struct{}{}

	inner = outer["a"]
	t.Log(len(inner))
}

func TestSliceMap(t *testing.T) {
	sliceMap := []map[string]int{
		{"1": 1},
		{"2": 2},
		{"3": 3},
		{"4": 4},
		{"5": 5},
		{"6": 6},
		{"7": 7},
	}

	sliceMapAnother := make([]map[string]int, 0, len(sliceMap))
	for _, m := range sliceMap {
		sliceMapAnother = append(sliceMapAnother, m)
	}

	t.Log(sliceMapAnother)
}

type MapValue struct {
	Name  string
	Age   int
	Value StructValue
}

// 从输出可以看出，这样并不能修改到MapValue的值
func TestUpdateMapValue(t *testing.T) {
	m := map[int]MapValue{
		1: {
			Name: "zhangsan",
			Age:  11,
		},
		2: {
			Name: "lisi",
			Age:  15,
		},
	}

	for k, v := range m {
		if k == 2 {
			v.Name = "wangmazi"
		}
	}

	t.Logf("%+v", m) // map[1:{Name:zhangsan Age:11} 2:{Name:lisi Age:15}]
}

func TestUpdateMapValue2(t *testing.T) {
	m := map[int]*MapValue{
		1: {
			Name: "zhangsan",
			Age:  11,
			Value: StructValue{
				IntValue:    11,
				StringValue: "zhangsan",
			},
		},
		2: {
			Name: "lisi",
			Age:  15,
			Value: StructValue{
				IntValue:    15,
				StringValue: "lisi",
			},
		},
	}

	for k, v := range m {
		if k == 2 {
			v.Name = "wangmazi"
			v.Value.StringValue = "wangmazi"
		}
	}

	for k, v := range m {
		t.Logf("k:%v, v:%v", k, v)
	}
}
