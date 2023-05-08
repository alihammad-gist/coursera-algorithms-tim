package main

import "fmt"

func main() {
	fmt.Println(BabageProduct(123345, 332))
	fmt.Println(MergeSort([]int{3, 4, 5, 1, 2, 3, 4, 1, 8, 1299, 343254, 345, 356345345, 345, 345, 34, 35, 345}))
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

// Merge sort
// 2. Merge
func MergeSort(xs []int) []int {
	// len
	l := len(xs)

	// Base case
	if l <= 1 {
		// single item array is already sorted. KAPISH!
		return xs
	}

	// 1. Split (time complexity: logn)
	as, bs := MergeSort(xs[l/2:]), MergeSort(xs[:l/2])

	// 2. Merge two sorted arrays as, bs (n: in linear time)
	sr := []int{}
	for i, j, k := 0, 0, 0; k < l; k++ {
		if i == len(as) {
			sr = append(sr, bs[j:]...)
			break
		}

		if j == len(bs) {
			sr = append(sr, as[i:]...)
			break
		}

		if as[i] < bs[j] {
			sr = append(sr, as[i])
			i++
		} else {
			sr = append(sr, bs[j])
			j++
		}
	}

	return sr
}
