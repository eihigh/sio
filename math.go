package sio

import (
	"math"
	"math/cmplx"
)

// UnitVector returns rotated unit vector
func UnitVector(angle float64) complex128 {
	return cmplx.Pow(complex(1, 0), complex(angle, 0))
}

// Rot returns (360/n)° rotated unit vector
func Rot(n float64) complex128 {
	return UnitVector(4.0 / n)
}

// Wave returns sin(2π*ratio)
func Wave(ratio, offset float64) float64 {
	return math.Sin(math.Pi*2*ratio + math.Pi*offset*2)
}

// UWave returns unsigned Wave, [0.0, 1.0]
func UWave(ratio, offset float64) float64 {
	return 0.5 + 0.5*Wave(ratio, offset)
}
