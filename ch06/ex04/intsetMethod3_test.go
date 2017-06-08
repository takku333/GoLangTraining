package main

import (
	"testing"
	"reflect"
)

func TestElms(t *testing.T){
	tests := []struct {
        input	[]int
    }{
		{
			input: []int{1, 2, 3, 63},
        },				
		{
			input: []int{1, 64, 256},
        },
    }
    for i, test := range tests {
		var s IntSet
        s.AddAll(test.input...)
		result := s.Elms()
		if !reflect.DeepEqual(test.input, result) {
            t.Errorf("test[%d] result:%d, want:%d\n", i, result, test.input)
        }
    }
}


