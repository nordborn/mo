// mo (think Monadic Operations)
// provides minimally sufficient yet powerful abstractions
// inspired by Haskell and Rust (names from Rust)
// for shorter and expressive programming.
// The package provides 2 data types Option and Result.
// and a helper function Catch.
// For those who worked with languages focused on null safety,
// Options' TryOr(default) is the point, because
// it provides the same semantic.
// For others Result can be an universal data type for
// expressive programming.
//
// To keep the things cleaner, there are no map-like methods/functions.
// Just extract value and process it or fail with Try.
//
// In any case, the key benefits like type-safe nil handling,
// shorter error handling are in place.
//
// Enjoy!
package mo
