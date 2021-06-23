package basic

import (
	"context"
	"testing"
)

type contextKey struct{}

func TestNilTypeAssertion(t *testing.T) {
	var k context.Context
	t.Log(k)
}

func TestTrueOrFalse(t *testing.T) {
	t.Log(1 == 0 || 2 == 2 && 1 == 1)
}
