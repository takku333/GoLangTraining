// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reser
package main

import (
	"testing"
)

func TestVariableJoin(t *testing.T){
	tests := []struct {
        input	[]string
        want 	string
    }{
        {
            input: []string{""},
			want: "",
        },
		{
            input: []string{"a"},
			want: "a",
        },
		{
            input: []string{"a", "b", "c"},
			want: "a b c",        
		},        
    }
    for i, test := range tests {
        result := variableJoin(" ", test.input...)
        if test.want != result {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}
