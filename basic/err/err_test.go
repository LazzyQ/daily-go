package main

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestWrapError(t *testing.T) {
	errLevel1 := fmt.Errorf("errLevel1")
	errLevel2 := fmt.Errorf("errLevel2 wrap err: %w", errLevel1)
	errLevel3 := fmt.Errorf("errLevel3 wrap err: %w", errLevel2)
	t.Log(errLevel3)

	unwrapErrLevel2 := errors.Unwrap(errLevel3)
	unwrapErrLevel1 := errors.Unwrap(unwrapErrLevel2)
	t.Log(unwrapErrLevel2 == errLevel2)
	t.Log(unwrapErrLevel1 == errLevel1)
}

func TestErrorsWrap(t *testing.T) {
	errLevel1 := fmt.Errorf("errLevel1")
	errLevel2 := errors.Wrap(errLevel1, "errLevel2 wrap")
	errLevel3 := errors.Wrap(errLevel2, "errLevel3 wrap")
	t.Log(errLevel3)
}
