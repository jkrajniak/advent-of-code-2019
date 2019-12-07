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
			Process(c.Tape, nil)
			assert.Equal(t, c.ExpectedValue, c.Tape[c.Ref])
		})
	}
}

func TestThrust(t *testing.T) {
	testCase := []struct {
		Tape   []int
		Phase  []int
		Output int
	}{{
		[]int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
		[]int{4, 3, 2, 1, 0},
		43210,
	},
		{
			[]int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23,
				101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			[]int{0, 1, 2, 3, 4},
			54321,
		},
		{
			[]int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33,
				1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			[]int{1, 0, 4, 3, 2},
			65210,
		},
	}

	for i, c := range testCase {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			thrust := ComputeThrust(c.Tape, c.Phase)
			assert.Equal(tt, c.Output, thrust)
		})
	}
}

func TestExtendThurst(t *testing.T) {
	testCase := []struct {
		Tape   []int
		Phase  []int
		Output int
	}{{
		[]int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,
			27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5},
		[]int{9,8,7,6,5},
		43210,
	}}

	for i, c := range testCase {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			thrust := ComputeThrustLoop(c.Tape, c.Phase)
			assert.Equal(tt, c.Output, thrust)
		})
	}
}
