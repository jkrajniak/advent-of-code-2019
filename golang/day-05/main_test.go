package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		Tape          []int
		Ref           int
		ExpectedValue int
	}{
		{
			[]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			0,
			3500,
		},
		{
			[]int{1, 0, 0, 0, 99},
			0,
			2,
		},
		{
			[]int{2, 3, 0, 3, 99},
			3,
			6,
		},
		{
			[]int{2, 4, 4, 5, 99, 0},
			5,
			9801,
		},
		{
			[]int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			0,
			30,
		},
		{
			[]int{4, 1, 4, 1, 4, 1, 99},
			0,
			4,
		},
		{
			[]int{1002, 4, 3, 4, 33},
			4,
			99,
		},
	}

	for i, c := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			Process(c.Tape)
			assert.Equal(tt, c.ExpectedValue, c.Tape[c.Ref])
		})
	}
}
