package basic

import (
	"encoding/json"
	"testing"
	"time"
)

func TestMarshalJSON(t *testing.T) {
	s := Student{
		Birthday: time.Now(),
	}

	str, _ := json.Marshal(s)
	t.Log(string(str))

	bytes := []byte(`{"birthday": }`)

	var s1 Student
	_ = json.Unmarshal(bytes, &s1)
	t.Log(s1)
}

func TestMarshalInt(t *testing.T) {
	bytes, _ := json.Marshal(1)
	t.Log(bytes)

	var n int
	_ = json.Unmarshal(bytes, &n)
	t.Log(n)
}
