package one

import (
	"testing"
)

func TestHighSchoolMult(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"mul#1", args{1235, 5678}, 7012330},
		{"mul#2", args{12, 5}, 60},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighSchoolMult(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("HighSchoolMult() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHighSchoolAddition(t *testing.T) {
	type args struct {
		n []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sum", args{[]int{1, 2, 33}}, 36},
		{"long_sum", args{[]int{11231, 123122, 309093}}, 443446},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HighSchoolAddition(tt.args.n); got != tt.want {
				t.Errorf("HighSchoolAddition() = %v, want %v", got, tt.want)
			}
		})
	}
}
