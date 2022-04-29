package main

import (
	"testing"
)

func TestGoroutineSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive numbers",
			args: args{[]int{3, 4, 5, 6}},
			want: 18,
		},
		{
			name: "Negative numbers",
			args: args{[]int{-3, -4, -5, -6}},
			want: -18,
		},
		{
			name: "Positive and negative numbers",
			args: args{[]int{3, 4, -5, 6}},
			want: 8,
		},
		{
			name: "Empty",
			args: args{[]int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GoroutineSum(tt.args.arr); got != tt.want {
				t.Errorf("GoroutineSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassicSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Positive numbers",
			args: args{[]int{3, 4, 5, 6}},
			want: 18,
		},
		{
			name: "Negative numbers",
			args: args{[]int{-3, -4, -5, -6}},
			want: -18,
		},
		{
			name: "Positive and negative numbers",
			args: args{[]int{3, 4, -5, 6}},
			want: 8,
		},
		{
			name: "Empty",
			args: args{[]int{}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClassicSum(tt.args.arr); got != tt.want {
				t.Errorf("ClassicSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

var arr5000 = generateRandomIntArr(5000)
var arr150000 = generateRandomIntArr(150000)

func BenchmarkGoroutineSumInput_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoroutineSum(arr5000)
	}
}

func BenchmarkClassicSumInput_5000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassicSum(arr5000)
	}
}

func BenchmarkGoroutineSumInput_150000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoroutineSum(arr150000)
	}
}

func BenchmarkClassicSumInput_150000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ClassicSum(arr150000)
	}
}
