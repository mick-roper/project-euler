package solution6

import "testing"

func Test_sumSquareDifference(t *testing.T) {
	type args struct {
		ceiling int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first 10 natural numbers returns 2640",
			args: args{ceiling: 10},
			want: 2640,
		},
		{
			name: "first 10 natural numbers returns 2640",
			args: args{ceiling: 20},
			want: 41230,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumSquareDifference(tt.args.ceiling); got != tt.want {
				t.Errorf("sumSquareDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
