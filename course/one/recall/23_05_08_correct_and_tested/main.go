package main

import "fmt"

func main() {
	as := []int{5, 2, 3, 0, 9, 1, 4, 5, 6, 7}

	fmt.Println(as, MergSort(as))
}

func BabageProduct(x, y int) int {
	// basecase
	if x < 10 && y < 10 {
		return x * y
	}

	m := ilen(max(x, y)) / 2

	a, b := x/pow(10, m), x%pow(10, m)
	c, d := y/pow(10, m), y%pow(10, m)

	ac := BabageProduct(a, c)
	ad := BabageProduct(a, d)
	bc := BabageProduct(b, c)
	bd := BabageProduct(b, d)

	return pow(10, m*2)*ac + pow(10, m)*(ad+bc) + bd
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func ilen(a int) int {
	n := 0
	for a > 0 {
		a /= 10
		n += 1
	}
	return n
}

func pow(n, e int) int {
	a := 1
	for ; e > 0; e -= 1 {
		a *= n
	}
	return a
}

// Divide and conqure MergeSort
func MergSort(xs []int) []int {
	n := len(xs)

	if n <= 1 {
		// one element array is already sorted.
		return xs
	}

	as := MergSort(xs[:n/2])
	bs := MergSort(xs[n/2:])
	res := make([]int, n) // result merged array

	// Given as and bs two sorted array
	// need to merge them in a way that
	// time complexity is linear O(n) and
	// the property of sorting is preserved
	j, k := 0, 0
	for i := 0; i < n; i++ {
		if j == len(as) {
			sliceAssign(i, n, res, bs[k:])
			break
		}

		if k == len(bs) {
			sliceAssign(i, n, res, as[j:])
			break
		}

		// Assign the smallest number first (ASC)
		// so the output can be compared with sort.Ints()
		if as[j] < bs[k] {
			res[i] = as[j]
			j += 1
		} else {
			res[i] = bs[k]
			k += 1
		}
	}

	return res
}

func sliceAssign(s, e int, dest, src []int) {
	for i, j := s, 0; i < e; i, j = i+1, j+1 {
		dest[i] = src[j]
	}
}
