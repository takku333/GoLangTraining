package main

import (
	"testing"
	"reflect"
)

func TestTopoSort(t *testing.T){
	tests := []struct {
        input	map[string][]string
        want 	[]string
    }{
        {
			input: map[string][]string{
				"courseA": {"courseB"},
				"courseB": {"courseC"},
		},
			want: []string{"courseC","courseB","courseA"},
        },
    }
    for i, test := range tests {
        result := topoSort(test.input)
        if !reflect.DeepEqual(test.want, result) {
            t.Errorf("test[%d] result:%s, want:%s\n", i, result, test.want)
        }
    }
}