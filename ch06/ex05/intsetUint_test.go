package main

import (
	"testing"
	"reflect"
)

func TestAdd(t *testing.T){
	tests := []struct {
        input	[]int
		target  int
        want 	[]uint
    }{
        {
			input: []int{},
			target: 1,
			want: []uint{0x2},
        },
		{
			input: []int{1},
			target: 1,
			want: []uint{0x2},
        },
		{
			input: []int{1},
			target: 64,
			want: []uint{0x2,0x1},
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
		s.Add(test.target)
		result := s.words
		if !reflect.DeepEqual(test.want, result) {
            t.Errorf("test[%d] result:%v, want:%v\n", i, result, test.want)
        }
    }
}

func TestHas(t *testing.T){
	tests := []struct {
        input	[]int
		target  int
        want 	bool
    }{
        {
			input: []int{},
			target: 1,
			want: false,
        },
		{
			input: []int{1, 2, 3, 63},
			target: 1,
			want: true,
        },
		{
			input: []int{1, 2, 3, 63},
			target: 4,
			want: false,
        },
		{
			input: []int{1, 2, 3, 63},
			target: 64,
			want: false,
        },						
		{
			input: []int{1, 64, 256},
			target: 64,
			want: true,
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
		result := s.Has(test.target)
		if test.want != result {
            t.Errorf("test[%d] result:%v, want:%v\n", i, result, test.want)
        }
    }
}

func TestString(t *testing.T){
	tests := []struct {
        input	[]int
        want 	string
    }{
        {
			input: []int{},
			want: "{}",
        },
		{
			input: []int{1, 2},
			want: "{1 2}",
        },
		{
			input: []int{1, 64, 256},
			want: "{1 64 256}",
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}

func TestUnion(t *testing.T){
	tests := []struct {
        inputS	[]int
		inputT	[]int
        want 	string
    }{
		{
			inputS: []int{1, 2, 64, 65},
			inputT: []int{1, 64, 256},			
			want: "{1 2 64 65 256}",
        },				
		{
			inputS: []int{1, 64, 256},	
			inputT: []int{1, 2, 64, 65},		
			want: "{1 2 64 65 256}",
        },
    }
    for i, test := range tests {
		var s, s2 IntSet
		for _, num := range test.inputS{
			s.Add(num)
		}
		for _, num := range test.inputT{
			s2.Add(num)
		}		
		s.UnionWith(&s2)
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}



