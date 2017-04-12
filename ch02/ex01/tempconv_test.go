// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package tempconv

import (
	"testing"
)

const (
    epsilon = 0.0001
)

func TestCToK(t *testing.T) {
    tests := []struct {
        c    Celsius
        want Kelvin
    }{
        {
            c: -273.15,
            want: 0,
        },
    }
    for i, test := range tests {
        result := CToK(test.c)
        if (test.want - result) > epsilon {
            t.Errorf("test[%d] result:%f, want:%f\n", i, result, test.want)
        }
    }
}

func TestFToK(t *testing.T) {
    tests := []struct {
        f    Fahrenheit
        want Kelvin
    }{
        {
            f: 100,
            want: 310.9278,
        },
    }
    for i, test := range tests {
        result := FToK(test.f)
        if (test.want - result) > epsilon {
            t.Errorf("test[%d] result:%f, want:%f\n", i, result, test.want)
        }
    }
}

func TestKToC(t *testing.T) {
    tests := []struct {
        k    Kelvin
        want Celsius
    }{
        {
            k:    0,
            want: -273.15,
        },
    }
    for i, test := range tests {
        result := KToC(test.k)
        if (test.want - result) > epsilon {
            t.Errorf("test[%d] result:%f, want:%f\n", i, result, test.want)
        }
    }
}

func TestKToF(t *testing.T) {
    tests := []struct {
        k    Kelvin
        want Fahrenheit
    }{
        {
            k: 310.9278,
            want: 100,
        },
    }
    for i, test := range tests {
        result := KToF(test.k)
        if (test.want - result) > epsilon {
            t.Errorf("test[%d] result:%f, want:%f\n", i, result, test.want)
        }
    }
}

func TestKevinString(t *testing.T) {
    tests := []struct {
        k    Kelvin
        want string
    }{
        {
            k: 0,
            want: "0K",
        },
    }
    for i, test := range tests {
        if test.want != test.k.String() {
            t.Errorf("test[%d] result:%s, want:%s\n", i, test.k.String(), test.want)
        }
    }
}