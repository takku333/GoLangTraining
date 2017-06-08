package main

import (
	"testing"
)

func TestIntersectWith(t *testing.T){
	tests := []struct {
        inputS	[]int
		inputT	[]int
        want 	string
    }{
         {
			inputS: []int{},
			inputT: []int{},
			want: "{}",
        },
		{
			inputS: []int{1, 2, 64, 65},
			inputT: []int{1, 64, 256},			
			want: "{1 64}",
        },				
		{
			inputS: []int{1, 64, 256},	
			inputT: []int{1, 2, 64, 65},		
			want: "{1 64}",
        },
    }
    for i, test := range tests {
		var s, s2 IntSet
        s.AddAll(test.inputS...)
		s2.AddAll(test.inputT...)
		s.IntersectWith(&s2)
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}

func TestDifferenceWith(t *testing.T){
	tests := []struct {
        inputS	[]int
		inputT	[]int
        want 	string
    }{
         {
			inputS: []int{},
			inputT: []int{},
			want: "{}",
        },
		{
			inputS: []int{1, 2, 64, 65},
			inputT: []int{1, 64, 256},			
			want: "{2 65}",
        },				
		{
			inputS: []int{1, 64, 256},	
			inputT: []int{1, 2, 64, 65},		
			want: "{256}",
        },
    }
    for i, test := range tests {
		var s, s2 IntSet
        s.AddAll(test.inputS...)
		s2.AddAll(test.inputT...)
		s.DifferenceWith(&s2)
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}

func TestSymmetricDifferenceWith(t *testing.T){
	tests := []struct {
        inputS	[]int
		inputT	[]int
        want 	string
    }{
         {
			inputS: []int{},
			inputT: []int{},
			want: "{}",
        },
		{
			inputS: []int{1, 2, 64, 65},
			inputT: []int{1, 64, 256},			
			want: "{2 65 256}",
        },				
		{
			inputS: []int{1, 64, 256},	
			inputT: []int{1, 2, 64, 65},		
			want: "{2 65 256}",
        },
    }
    for i, test := range tests {
		var s, s2 IntSet
        s.AddAll(test.inputS...)
		s2.AddAll(test.inputT...)
		s.SymmetricDifferenceWith(&s2)
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}