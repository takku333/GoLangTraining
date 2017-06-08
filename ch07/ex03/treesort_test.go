package main

import (
	"testing"
	"fmt"
)

func TestString(t *testing.T){
	tests := []struct {
        input	[]int
        want 	string
    }{
        {
			input: []int{},
			want: "[]",
        },
		{
			input: []int{1},
			want: "[1]",
        },
		{
			input: []int{1, 2, 3},
			want: "[1 2 3]",
        },
    }
    for i, test := range tests {
		var tr *tree
		for _, num := range test.input {
			tr = add(tr, num)
		}
		result := fmt.Sprint(tr.String())
		if test.want != result {
            t.Errorf("test[%d] result:%v, want:%v\n", i, result, test.want)
        }
    }
}
