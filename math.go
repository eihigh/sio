package sio

import (
	"math"
	"math/cmplx"
)

// AbsSq returns abs*abs, more faster
func AbsSq(c complex128) float64 {
	x, y := real(c), imag(c)
	return x*x + y*y
}

// UnitVector returns rotated unit vector
func UnitVector(angle float64) complex128 {
	return cmplx.Pow(0+1i, complex(angle, 0))
}

// Rot returns (360/n)° rotated unit vector
func Rot(n int) complex128 {
	return UnitVector(4.0 / float64(n))
}

// Wave returns sin(2π*ratio)
func Wave(ratio float64) float64 {
	return math.Sin(math.Pi * 2 * ratio)
}

// UWave returns unsigned Wave, [0.0, 1.0]
func UWave(ratio float64) float64 {
	return 0.5 + 0.5*Wave(ratio)
}
