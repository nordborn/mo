package mo

import "fmt"

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
//		 func someFunc(...) (ret Result[int]) {
//		 	errorContext := "someFunc"
//			defer mo.Catch(errorContext, &ret)
//		 	...
//		    1 / 0
//	     	...
//			return
//		 }
func Catch[T any](on string, ret *Result[T]) {
	if r := recover(); r != nil {
		if err, ok := r.(error); ok {
			*ret = Err[T](err).On(on)
			return
		}
		*ret = Err[T](fmt.Errorf("catched panic: %v", r)).On(on)
	}
}
