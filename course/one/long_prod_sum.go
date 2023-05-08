package one

import (
	"math"
)

// HighSchoolMult implements the standard algorithm everyone uses
// on paper to calculate product of two numbers.
func HighSchoolMult(x, y int) int {
	xs := getDigits(x) // spliting digits
	ys := getDigits(y)
	xl := len(xs)
	yl := len(ys)
	in := make([]int, xl) // intermedite number that will be summed

	for i := 0; i < xl; i++ {
		a, b, c, n := 0, 0, 0, 0

		for j := 0; j < yl; j++ {
			// reading the list from opposite direction
			// ie. rtl multiplication of arabic numerals
			a = xs[xl-i-1]
			b = ys[yl-j-1]
			n, c = mulSingleDigit(a, b, c)
			in[i] += n * int(math.Pow(10, float64(i+j))) // joining digits
		}

		// the last carry.
		in[i] += c * int(math.Pow(10, float64(i+yl)))
	}

	return HighSchoolAddition(in)
}

func HighSchoolAddition(n []int) int {
	mx := struct {
		val int
		idx int
		dl  int
	}{
		val: n[0],
		idx: 0,
		dl:  len(getDigits(n[0])),
	}

	// converting numbers to digits
	ns := [][]int{}
	for i, x := range n {
		dgs := getDigits(x)
		ns = append(ns, dgs)
		if x > mx.val {
			mx.val = x
			mx.idx = i
			mx.dl = len(dgs)
		}
	}

	// preparing numbers to before summation,
	// prepending zeros so addition becomes easier
	for i := 0; i < len(ns); i++ {
		ns[i] = padZeros(ns[i], mx.dl-len(ns[i]))
	}

	cr := 0 // carry
	sum := 0
	for i := 0; i < mx.dl; i++ {
		col := []int{}
		ri := mx.dl - i - 1
		s := 0
		for j := 0; j < len(n); j++ {
			// rtl addition
			col = append(col, ns[j][ri])
		}
		cr, s = addSingleDigits(cr, col...)
		sum += (s * int(math.Pow(10, float64(i))))
	}

	return sum
}

// mulSingleDigit returns product, carry
func mulSingleDigit(a, b, c int) (int, int) {
	if a > 9 || b > 9 || c > 9 {
		panic("mulSingleDigit only accepts single digits")
	}

	n := getDigits((a * b) + c)

	if len(n) == 2 {
		return n[1], n[0]
	}

	return n[0], 0
}

// addSingleDigits returns carry, sum
func addSingleDigits(c int, dgs ...int) (int, int) {
	sum := c
	for _, n := range dgs {
		sum += n
	}

	return sum / 10, sum % 10
}

// padZeros prepends n zeros to the pass
// in slice
func padZeros(xs []int, n int) []int {
	ns := make([]int, n)
	return append(ns, xs...)
}
