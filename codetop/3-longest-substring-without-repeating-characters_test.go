package codetop

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test-1",
			args: args{
				s: " ",
			},
			want: 1,
		},
		{
			name: "test-2",
			args: args{
				s: "aa",
			},
			want: 1,
		},
		{
			name: "test-3",
			args: args{
				s: "abc",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
