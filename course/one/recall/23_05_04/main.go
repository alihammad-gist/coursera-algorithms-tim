package main

import "fmt"

func main() {
	// fmt.Println(pow(10, 3))
	// fmt.Println(max(13, 3))
	// fmt.Println(ilen(3333))
	n := KaratsubaMul(22, 278)
	fmt.Println("Ans:", n, "Expected", 22*278)
}

func KaratsubaMul(x, y int) int {
	if x < 10 && y < 10 {
		return x * y
	}

	n := ilen(max(x, y))

	a := x / pow(10, n/2)
	b := x % pow(10, n/2)
	c := y / pow(10, n/2)
	d := y % pow(10, n/2)

	fmt.Println("X:", x, a, b)
	fmt.Println("Y:", y, c, d)
	ac := KaratsubaMul(a, c)
	bd := KaratsubaMul(b, d)
	e := KaratsubaMul(a+b, c+d) - ac - bd

	// fmt.Println(x, a, b)
	// fmt.Println(y, c, d)
	fmt.Println((pow(10, n) * ac) + (pow(10, n/2) * e) + bd)

	return (pow(10, n) * ac) + (pow(10, n/2) * e) + bd
}

func pow(n, exp int) int {
	a := 1
	for exp > 0 {
		a *= n
		exp -= 1
	}
	return a
}

func ilen(x int) int {
	n := 0
	for x > 0 {
		x /= 10
		n += 1
	}
	return n
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
