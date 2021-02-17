package problem1

import "testing"

func Test_sumOfMultiplesOf3And5(t *testing.T) {
	type args struct {
		ceiling int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "ceiling of 10 returns 23",
			args: args{ceiling: 10},
			want: 23,
		},
		{
			name: "ceiling of 1000 returns 23",
			args: args{ceiling: 1000},
			want: 233168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfMultiplesOf3And5(tt.args.ceiling); got != tt.want {
				t.Errorf("sumOfMultiplesOf3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
