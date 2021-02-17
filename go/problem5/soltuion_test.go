package problem5

import "testing"

func Test_smallestMultiple(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "smallest number that can be divided between 1 to 10 without remainder is 2520",
			args: args{start: 1, end: 10},
			want: 2520,
		},
		{
			name: "smallest number that can be divided between 1 to 10 without remainder is 2520",
			args: args{start: 1, end: 20},
			want: 2520,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestMultiple(tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("smallestMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
