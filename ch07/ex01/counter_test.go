package main

import (
	"testing"
	"fmt"
)

func TestWordCounter(t *testing.T){
	tests := []struct {
        input	string
        want 	string
    }{
        {
			input: "",
			want: "0",
        },
		{
			input: "a",
			want: "1",
        },
		{
			input: "a b c abc",
			want: "4",
        },
    }
    for i, test := range tests {
		var wc WordCounter
		wc.Write([]byte(test.input))
		result := fmt.Sprint(wc)
		if test.want != result {
            t.Errorf("test[%d] result:%v, want:%v\n", i, result, test.want)
        }
    }
}

func TestLineCounter(t *testing.T){
	tests := []struct {
        input	string
        want 	string
    }{
        {
			input: "",
			want: "0",
        },
		{
			input: "a",
			want: "1",
        },
		{
			input: "a\nb\nc\nabc",
			want: "4",
        },
    }
    for i, test := range tests {
		var lc LineCounter
		lc.Write([]byte(test.input))
		result := fmt.Sprint(lc)
		if test.want != result {
            t.Errorf("test[%d] result:%v, want:%v\n", i, result, test.want)
        }
    }
}
