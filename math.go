package main

import "math"

// Abs - Go's math lib only has a f64 version for some reason.
func Abs(x float32) float32 {
	return math.Float32frombits(math.Float32bits(x) &^ (1 << 31))
}
