package dfs

import (
	"reflect"
	"testing"
)

func Test_findOrder210(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			want: []int{0, 1},
		},
		{
			name: "2",
			args: args{
				numCourses: 4,
				prerequisites: [][]int{
					{1, 0},
					{2, 0},
					{3, 1},
					{3, 2},
				},
			},
			want: []int{0,2,1,3},
		},
		{
			name: "3",
			args: args{
				numCourses: 2,
				prerequisites: [][]int{
					{0, 1},
					{1, 0},
				},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrder210(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrder210() = %v, want %v", got, tt.want)
			}
		})
	}
}
