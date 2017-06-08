package main

import (
	"testing"
)

func TestAddAll(t *testing.T){
	tests := []struct {
        input	[]int
        want 	string
    }{
         {
			input: []int{},
			want: "{}",
        },
		{
			input: []int{1, 2, 3, 63},
			want: "{1 2 3 63}",
        },				
		{
			input: []int{1, 64, 256},
			want: "{1 64 256}",
        },
    }
    for i, test := range tests {
		var s IntSet
        s.AddAll(test.input...)
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}


