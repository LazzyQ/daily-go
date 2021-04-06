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
