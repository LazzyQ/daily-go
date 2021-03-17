package basic

import (
	"testing"

	"golang.org/x/text/unicode/norm"
)

func TestNFKD(t *testing.T) {

	normTest := []struct {
		Input  string
		Expect string
	}{
		{"我是四川的，你呢？", "我是四川的,你呢?"},
		{"I'm from sichuan, and you？", "I'm from sichuan, and you?"},
	}

	for _, tt := range normTest {
		actual := norm.NFKD.String(tt.Input)
		t.Log(actual)
		if tt.Expect != actual {
			t.Fatalf("norm.NFKD.String(%s) expect is %v, but actual is %v", tt.Input, tt.Expect, actual)
		}
	}
}
