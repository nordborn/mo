package mo_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nordborn/mo"
	"github.com/stretchr/testify/assert"
)

func TestCatch(t *testing.T) {
	res := divideResultCatched(10, 0)
	fmt.Printf("%+v\n", res)
	assert.Equal(t, -1, res.TryOr(-1))
}

func divideResultCatched(a, b int) (ret mo.Result[int]) {
	defer mo.Catch("divideResultCatched", &ret)
	ret = mo.Ok(a / b)
	return
}

func TestCatchTry(t *testing.T) {
	res := process()
	fmt.Printf("%+v\n", res)
	assert.Equal(t, false, res.IsOk())
}

func process() (ret mo.Result[int]) {
	defer mo.Catch("process", &ret)

	file := mo.ResultFrom(openFile()).Try("open file")
	data := readFile(file).Try()
	ret = parse(data)
	return
}

// Mock functions to simulate operations
func openFile() (*File, error) {
	return nil, errors.New("file not found")
}

func readFile(_ *File) mo.Result[[]byte] {
	return mo.Ok([]byte("dummy data"))
}

func parse(_ []byte) mo.Result[int] {
	return mo.Ok(42)
}

type File struct{}

func TestCatchErr(t *testing.T) {
	ret := processErr()
	assert.Equal(t, false, ret.IsOk())
}

func processErr() (ret mo.Result[int]) {
	defer mo.Catch("processErr", &ret)
	mo.TryErr(retErr())
	return
}

func retErr() error {
	return errors.New("a test error")
}

func BenchmarkPureDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		defer func() {}()
	}
}

func BenchmarkErr(b *testing.B) {
	err0 := errors.New("some error")
	for i := 0; i < b.N; i++ {
		func() (err error) {
			ctx := "some func"
			err = fmt.Errorf("%s: %d: %w", ctx, i, err0)
			return
		}()
	}
}

func BenchmarkCatch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() (ret mo.Result[int]) {
			defer mo.Catch("some func", &ret)
			panic(i)
		}()
	}
}
