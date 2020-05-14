// Package triangle is Exercism.io side exercise
package triangle

import "math"

type Kind int

const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
	Deg        // degenerate
)

// KindFromSides — determine if a triangle is equilateral, isosceles, scalene or degenerate
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isTriangle(a, b, c):
		return NaT
	case a == b && b == c && a == c:
		return Equ
	case a+b == c || a+c == b || b+c == a:
		return Deg
	case a != b && b != c && a != c:
		return Sca
	case (a == b) || (b == c) || (a == c):
		return Iso
	default:
		return NaT
	}
}

func isTriangle(a, b, c float64) bool {
	switch {
	case math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c):
		return false
	case math.IsInf(a, 1) || math.IsInf(a, -1) || math.IsInf(b, 1) || math.IsInf(b, -1) || math.IsInf(c, 1) || math.IsInf(c, -1):
		return false
	case a <= 0 || b <= 0 || c <= 0:
		return false
	case a+b < c || a+c < b || b+c < a:
		return false
	default:
		return true
	}
}
