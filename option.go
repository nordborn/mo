package mo

import (
	"encoding/json"
	"errors"
)

// Option type implementation with JSON marchalling/unmarshalling support
type Option[T any] struct {
	val    T
	isSome bool
}

// Some is a constructor of some Option (valuable data)
func Some[T any](v T) Option[T] {
	return Option[T]{val: v, isSome: true}
}

// None is a constructor of none Option (nil data)
func None[T any]() Option[T] {
	return Option[T]{}
}

// OptionFrom makes a new Option from `val, ok` statement
func OptionFrom[T any](v T, ok bool) Option[T] {
	if !ok {
		return None[T]()
	}
	return Some(v)
}

// IsSome checks is some (true) or none (false)
func (o Option[T]) IsSome() bool {
	return o.isSome
}

// Try unwraps or panics (can be intercepted by `Catch`)
func (o Option[T]) Try() T {
	if !o.isSome {
		panic(errors.New("try none"))
	}
	return o.val
}

// TryOr extrats value or provides defaultTry
func (o Option[T]) TryOr(defaultTry T) T {
	if o.isSome {
		return o.val
	}
	return defaultTry
}

// Unpack converts Option to `val, ok`
func (o Option[T]) Unpack() (T, bool) {
	return o.val, o.isSome
}

// JSON Marshaling implementation
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.isSome {
		return json.Marshal(o.val)
	}
	return json.Marshal(nil)
}

// JSON Unmarshaling implementation
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	var raw json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	if raw == nil || (len(raw) == 4 && string(raw) == "null") {
		o.isSome = false
		o.val = *new(T) // Set to zero value of T
		return nil
	}

	var value T
	if err := json.Unmarshal(raw, &value); err != nil {
		return err
	}

	o.val = value
	o.isSome = true
	return nil
}
