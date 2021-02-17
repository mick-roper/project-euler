package problem3

import (
	"testing"
)

func Test_largestPrimeFactor(t *testing.T) {
	type args struct {
		prime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "largest prime factor of 17 is 17",
			args: args{prime: 17},
			want: 17,
		},
		{
			name: "largest prime factor of 12 is 3",
			args: args{prime: 12},
			want: 3,
		},
		{
			name: "largest prime factor of 147 is 7",
			args: args{prime: 147},
			want: 7,
		},
		{
			name: "largest prime factor of 13195 is 29",
			args: args{prime: 13195},
			want: 29,
		},
		{
			name: "largest prime factor of 600851475143 is 6857",
			args: args{prime: 600851475143},
			want: 6857,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPrimeFactor(tt.args.prime); got != tt.want {
				t.Errorf("largestPrimeFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "2 should be prime",
			args: args{n: 2},
			want: true,
		},
		{
			name: "7 should be prime",
			args: args{n: 7},
			want: true,
		},
		{
			name: "89 should be prime",
			args: args{n: 89},
			want: true,
		},
		{
			name: "10 should not be prime",
			args: args{n: 10},
			want: false,
		},
		{
			name: "99 should not be prime",
			args: args{n: 99},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.n); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 returns 3",
			args: args{n: 2},
			want: 3,
		},
		{
			name: "13 returns 17",
			args: args{n: 13},
			want: 17,
		},
		{
			name: "89 returns 97",
			args: args{n: 89},
			want: 97,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPrime(tt.args.n); got != tt.want {
				t.Errorf("nextPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
