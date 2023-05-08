package one

func BabageProd(x, y int) int {
	// base case
	if x < 10 && y < 10 {
		return x * y
	}
	n := maxLen(x, y)

	// number spliting
	a := x / pow(10, n/2)
	b := x % pow(10, n/2)

	c := y / pow(10, n/2)
	d := y % pow(10, n/2)

	// products (recursion)
	ac := BabageProd(a, c)
	ad := BabageProd(a, d)
	bc := BabageProd(b, c)
	bd := BabageProd(b, d)

	return pow(10, n)*ac + pow(10, n/2)*(ad+bc) + bd
}

func KaratsubaProd(x, y int) int {
	// karatsuba optimizes Babage's algorithm
	// by reducing a product operation at each
	// setp of recursion

	// base case
	if x < 10 && y < 10 {
		return x * y
	}
	n := maxLen(x, y)

	// number spliting
	a := x / pow(10, n/2)
	b := x % pow(10, n/2)

	c := y / pow(10, n/2)
	d := y % pow(10, n/2)

	// products (recursion)
	ac := KaratsubaProd(a, c)
	e := KaratsubaProd(a+b, c+d)
	bd := KaratsubaProd(b, d)

	return pow(10, n)*ac + bd + pow(10, n/2)*(e-ac-bd)
}

func nlen(n int) int {
	c := 0
	for n > 0 {
		c += 1
		n = n / 10
	}
	return c
}

func maxLen(a, b int) int {
	al, bl := nlen(a), nlen(b)
	if al > bl {
		return al
	}

	return bl
}

func pow(n, p int) int {
	a := n
	p -= 1
	for p > 0 {
		p -= 1
		a *= n
	}

	return a
}
