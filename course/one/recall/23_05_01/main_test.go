package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	t.Log(
		MergeSort([]int{4, 5, 7, 1, 3, 10, 8, 9, 2, 6}),
	)
	assert.Equal(
		t,
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		MergeSort([]int{4, 5, 7, 1, 3, 10, 8, 9, 2, 6}),
	)
}

// BabageProduct emloys the karatsuba algorithm as follows
//
// # Components of x and y
//
// n is the either the length of x or y (whichever is longer)
// x = 10^(n/2) * a + b
// y = 10^(n/2) * c + d
//
//
// # Recursive formula to calculate product
//
// x.y = 10^n ac + 10^(n/2) (ad + bc) + bd
func BabageProduct(x, y int) int {
	if x < 9 && y < 9 {
		return x * y
	}

	n := ilen(max(x, y))

	a := x / pow(10, n/2)
	b := x % pow(10, n/2)

	c := y / pow(10, n/2)
	d := y % pow(10, n/2)
	fmt.Println(x, a, b, "-", y, c, d)

	// products
	ac := BabageProduct(a, c)
	ad := BabageProduct(a, d)
	bc := BabageProduct(b, c)
	bd := BabageProduct(b, d)

	return pow(10, n)*ac + pow(10, n/2)*(ad+bc) + bd
}

func ilen(n int) int {
	l := 0
	for n > 0 {
		n /= 10
		l++
	}

	return l
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pow(n, p int) int {
	a := 1
	for p > 0 {
		a *= n
		p -= 1
	}
	return a
}

func MergeSort(xs []int) []int {
	n := len(xs)

	// base
	if n < 2 {
		return xs
	}

	// Splits
	as, bs := MergeSort(xs[:n/2]), MergeSort(xs[n/2:])

	// Merge
	res := make([]int, 0)
	for i, j, k := 0, 0, 0; i < n; i++ {
		if j >= len(as) {
			// as has exhausted
			res = append(res, bs[k:]...)
			break
		}

		if k >= len(bs) {
			// bs has exhausted
			res = append(res, as[j:]...)
			break
		}

		if as[j] < bs[k] {
			res = append(res, as[j])
			j++
		} else {
			res = append(res, bs[k])
			k++
		}
	}

	return res
}
