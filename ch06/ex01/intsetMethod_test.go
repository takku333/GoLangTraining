package main

import (
	"testing"
)

func TestLen(t *testing.T){
	tests := []struct {
        input	[]int
        want 	int
    }{
        {
			input: []int{},
			want: 0,
        },
		{
			input: []int{1, 2, 3, 63},
			want: 4,
        },
		{
			input: []int{1, 64, 256},
			want: 3,
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
        result := s.Len()
		if test.want != result {
            t.Errorf("test[%d] result:%d, want:%d\n", i, result, test.want)
        }
    }
}

func TestRemove(t *testing.T){
	tests := []struct {
        input	[]int
		target  int
        want 	string
    }{
        {
			input: []int{},
			target: 1,
			want: "{}",
        },
		{
			input: []int{1, 2, 3, 63},
			target: 1,
			want: "{2 3 63}",
        },
		{
			input: []int{1, 2, 3, 63},
			target: 4,
			want: "{1 2 3 63}",
        },
		{
			input: []int{1, 2, 3, 63},
			target: 64,
			want: "{1 2 3 63}",
        },						
		{
			input: []int{1, 64, 256},
			target: 64,
			want: "{1 256}",
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
        s.Remove(test.target)
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}

func TestClear(t *testing.T){
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
			want: "{}",
        },
		{
			input: []int{1, 64, 256},
			want: "{}",
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
        s.Clear()
		result := s.String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}

func TestCopy(t *testing.T){
	tests := []struct {
        input	[]int
        want 	string
    }{
        {
			input: []int{},
        },
		{
			input: []int{1, 2, 3, 63},
        },
		{
			input: []int{1, 64, 256},
        },
    }
    for i, test := range tests {
		var s IntSet
		for _, num := range test.input{
			s.Add(num)
		}
		test.want = s.String()
        result := s.Copy().String()
		if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}

