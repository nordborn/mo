package mo

import "fmt"

// TryErr just panics on err (of course to be intercepted by `Catch`)
func TryErr(err error, on ...string) {
	if err != nil {
		if on == nil {
			panic(fmt.Errorf("try err: %w", err))
		} else {
			panic(fmt.Errorf("%s: try err: %w", on[0], err))
		}

	}
}
