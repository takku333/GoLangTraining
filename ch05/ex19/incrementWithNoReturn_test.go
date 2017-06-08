// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reser
package main

import (
	"testing"
)

func TestIncrement(t *testing.T){
	tests := []struct {
        input	int
        want 	int
    }{
        {
			want: 1,
        },
		{
            input: 1,
			want: 2,
        },
    }
    for i, test := range tests {
        result := increment(test.input)
        if test.want != result {
            t.Errorf("test[%d] result:%d, want:%d\n", i, result, test.want)
        }
    }
}
