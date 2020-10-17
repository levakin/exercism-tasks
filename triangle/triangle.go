// Package triangle implements functionality to determine kind of triangle
package triangle

import "math"

type Kind string

const (
	NaT Kind = "not a triangle"
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
)

// KindFromSides determines kind of triangle
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case !isTriangle(a, b, c):
		k = NaT
	case a == b && a == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	default:
		k = Sca
	}
	return k
}

func isTriangle(a, b, c float64) bool {
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return false
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return false
	}
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b < c || a+c < b || b+c < a {
		return false
	}
	return true
}
