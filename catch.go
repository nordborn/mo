package mo

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/nordborn/go-errow"
)

// Catch recovers Try panics and converts them to Err,
// as well as any runtime panics.
// It's useful to handle Try() calls and sets ret to Err[T].
// Note, that Try() panics on None or Err, so, intentionally,
// it should be used in cases when the code expects Some or Ok.
// Generally, errors are not a part of common code executions,
// thus, this approach doesn't generates overhead.
// Meanwhile, when you expect a lot of None or Err during unwrapping,
// better handle them by checking IsSome and isOk,
// but, the code becomes close to regular go code.
// Hint for Option, better use TryOr for defaults instead of Try,
// it rules too. See code examples and tests, or just try.
//
// Example:
//
//		 // on catch will return Err with "module.someFunc(arg1, arg2): <file/path.go:33>: runtime error: integer divide by zero"
//		 func someFunc(arg1, arg2) (ret Result[int]) {
//			defer mo.Catch(&ret, arg1, arg2)
//		 	...
//		    1 / 0
//	     	...
//			return ret
//		 }
func Catch[T any](ret *Result[T], on ...any) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			// WrapSkip(6, err) points to the place where the panic occured
			*ret = Err[T](errow.WrapSkip(6, err)).On(getOuterFuncName(), "(", fmt.Sprint(on...), ")")
			return
		}
		*ret = Err[T](fmt.Errorf("catched panic: %v", r)).On(getOuterFuncName(), "(", fmt.Sprint(on...), ")")
	}
}

// getOuterFuncName returns module.(receiver).func where Catch was deferred
func getOuterFuncName() string {
	pc, _, _, _ := runtime.Caller(4)
	fName := runtime.FuncForPC(pc).Name()
	s := strings.Split(fName, "/")
	// drop module path
	return s[len(s)-1]
}
