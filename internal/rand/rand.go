package rand

import "math/rand"

// IntRange returns, as an int, a non-negative pseudo-random number in the
// interval [min,max).
func IntRange(min, max int) int {
	return rand.Intn(max-min) + min
}
