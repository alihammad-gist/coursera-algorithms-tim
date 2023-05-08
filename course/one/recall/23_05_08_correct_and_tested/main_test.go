package main

import (
	"sort"
	"testing"

	"github.com/flyingmutant/rapid"
	"github.com/stretchr/testify/assert"
)

func TestBabageProduct(t *testing.T) {
	r := rapid.IntRange(1, 10000)

	rapid.Check(t, func(t *rapid.T) {
		x := r.Draw(t, "x")
		y := r.Draw(t, "y")

		assert.Equal(t,
			x*y,
			BabageProduct(x, y),
		)
	})
}

func TestMergeSort(t *testing.T) {
	gen := rapid.SliceOf(rapid.Int())

	rapid.Check(t, func(t *rapid.T) {
		sl := gen.Draw(t, "Merge sort test")
		ms := MergSort(sl)
		ss := make([]int, len(sl))
		copy(ss, sl)
		sort.Ints(ss)
		assert.Equal(t, ms, ss)
	})
}
