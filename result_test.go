package mo_test

import (
	"errors"
	"mo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	assert.Equal(t, 5, divideSafe(10, 2).TryOr(0))
	assert.Equal(t, -1, divideSafe(10, 0).TryOr(-1))
}

func divideSafe(a, b int) mo.Result[int] {
	if b == 0 {
		return mo.Err[int](errors.New("division by zero")).On("divideSafe")
	}
	return mo.Ok(a / b)
}
