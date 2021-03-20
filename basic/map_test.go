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
