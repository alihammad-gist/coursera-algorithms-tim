package one

func getDigits(x int) []int {
	xs := make([]int, 0)
	for x > 0 {
		xs = append([]int{x % 10}, xs...)
		x = x / 10
	}

	return xs
}
