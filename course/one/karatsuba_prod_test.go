package one

import (
	"testing"
)

func TestBabageProd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// single digit simplest
		{"babage_product#1", args{3, 2}, 6},
		{"babage_product#2", args{321, 2}, 642},
		{"babage_product#3", args{321214, 2}, 642428},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BabageProd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("BabageProd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKaratsubaProd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Karatsuba product#1", args{3, 2}, 6},
		{"Karatsuba product#2", args{321, 2}, 642},
		{"Karatsuba product#3", args{321214, 2}, 642428},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KaratsubaProd(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("KaratsubaProd() = %v, want %v", got, tt.want)
			}
		})
	}
}
