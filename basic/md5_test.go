package basic

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"testing"
)

func BenchmarkStringEqual(b *testing.B) {
	str := "关键词xxxxxxxxxxxxxxxxxxxx"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = str == "关键词xxxxxxxxxxxxxxxxxxxx"
	}
}

func BenchmarkMd5Equal(b *testing.B) {
	str := "关键词xxxxxxxxxxxxxxxxxxxx"
	h := md5.New()
	h.Write([]byte(str))
	var num int64
	data := h.Sum(nil)
	buf := bytes.NewBuffer(data)
	_ = binary.Read(buf, binary.LittleEndian, &num)
	//fmt.Println(num)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = -758800610277284536 == num
	}
}

