// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reser

package main

import (
	"testing"
)

func TestMax(t *testing.T){
	tests := []struct {
        input	[]int
        want 	int
    }{
        {
            input: []int{3},
			want: 3,
        },
		{
            input: []int{},
			want: 0,
        },
		{
            input: []int{1,2,3,4,5,5,4,3,2,1},
			want: 5,
        },        
    }
    for i, test := range tests {
        result := max(test.input...)
        if test.want != result {
            t.Errorf("test[%d] result:%d, want:%d\n", i, result, test.want)
        }
    }
}

func TestMin(t *testing.T){
	tests := []struct {
        input	[]int
        want 	int
    }{
        {
            input: []int{3},
			want: 3,
        },
		{
            input: []int{},
			want: 0,
        },
		{
            input: []int{1,2,3,4,5,5,4,3,2,1},
			want: 5,
        },        
    }
    for i, test := range tests {
        result := max(test.input...)
        if test.want != result {
            t.Errorf("test[%d] result:%d, want:%d\n", i, result, test.want)
        }
    }
}