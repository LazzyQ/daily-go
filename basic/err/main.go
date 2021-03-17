package main

import (
	"fmt"

	"github.com/pkg/errors"
)

var TEST_ERROR = fmt.Errorf("这是一个错误")

func stdError() error {
	return fmt.Errorf("发生了错误: %w", TEST_ERROR)
}

func errorsError() error {
	return errors.Wrapf(TEST_ERROR, "发生了错误")
}

func main() {
	err := stdError()
	fmt.Printf("stdError: %v\n", err)
	fmt.Printf("stdError: %+v\n", err)

	err = errorsError()
	fmt.Printf("errorsError: %v\n", err)
	fmt.Printf("errorsError: %+v\n", err)
}
