package main

import (
	"testing"

	"github.com/flyingmutant/rapid"
	"github.com/stretchr/testify/assert"
)

func TestKaratsuba(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		r := rapid.IntRange(1, 1234)
		x := r.Draw(t, "x")
		y := r.Draw(t, "y")

		assert.Equal(t, x*y, KaratsubaMul(x, y))
	})
}

// func FuzzKaratsuba(f *testing.F) {
// 	f.Fuzz(func(t *testing.T, x, y int) {
// 		if !overflows(x, y) && gtzero(x, y) {
// 			m := KaratsubaMul(x, y)
// 			assert.Equal(t, x*y, m)
// 		}
// 	})
// }
//
// func overflows(x, y int) bool {
// 	n := int64(x) * int64(y)
// 	if n == int64(x*y) {
// 		return true
// 	}
//
// 	return false
// }
//
// func gtzero(a, b int) bool {
// 	return a > 0 && b > 0
// }
