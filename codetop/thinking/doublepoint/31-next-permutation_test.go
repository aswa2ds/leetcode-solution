package doublepoint

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{2, 3, 1},
			},
			want: []int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
