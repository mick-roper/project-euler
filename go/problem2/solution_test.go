package problem2

import "testing"

func Test_fibonacciSum(t *testing.T) {
	type args struct {
		ceiling int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ceiling of 100, even fibs sum = 44",
			args: args{ceiling: 100},
			want: 44,
		},
		{
			name: "ceiling of 4000000, even fibs sum = 44",
			args: args{ceiling: 4000000},
			want: 4613732,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciSum(tt.args.ceiling); got != tt.want {
				t.Errorf("fibonacciSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
