package basic

import (
	"sync"
	"testing"
)

func BenchmarkSlicePool(b *testing.B) {

	p := sync.Pool{
		New: func() interface{} {
			s := make([]byte, 128)
			return s
		},
	}
	for i := 0; i < b.N; i++ {
		s := p.Get()
		if bytes, ok := s.([]byte); ok {
			bytes[0] = 0
		}
		p.Put(s)
	}
}

func BenchmarkPointSlicePool(b *testing.B) {
	p := sync.Pool{
		New: func() interface{} {
			s := make([]byte, 128)
			return &s
		},
	}
	for i := 0; i < b.N; i++ {
		s := p.Get()
		if bytes, ok := s.(*[]byte); ok {
			(*bytes)[0] = 0
		}
		p.Put(s)
	}
}
