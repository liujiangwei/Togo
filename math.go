package gophp

import (
	"math"
	"math/cmplx"
)

func Abs(n float64) float64 {
	return math.Abs(n)
}

func Acos(val complex128) complex128 {
	return cmplx.Acos(val)
}

func Acosh(val complex128) complex128 {
	return cmplx.Acosh(val)
}