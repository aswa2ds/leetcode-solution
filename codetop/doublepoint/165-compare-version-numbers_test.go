package doublepoint

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				version1: "1.0.1",
				version2: "1.0.0",
			},
			want: 1,
		},
		{
			name: "test2",
			args: args{
				version1: "1.0.1",
				version2: "1.0",
			},
			want: 1,
		},
		{
			name: "test3",
			args: args{
				version1: "1.0.0.0",
				version2: "1.0",
			},
			want: 0,
		},
		{
			name: "test4",
			args: args{
				version1: "0.1",
				version2: "1.1",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
