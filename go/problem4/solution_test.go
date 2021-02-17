package problem4

import "testing"

func Test_generatePalindrome(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "2 digit numbers returns 9009",
			args: args{digits: 2},
			want: 9009,
		},
		{
			name: "3 digit numbers returns 580085",
			args: args{digits: 3},
			want: 906609,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generatePalindrome(tt.args.digits); got != tt.want {
				t.Errorf("generatePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
