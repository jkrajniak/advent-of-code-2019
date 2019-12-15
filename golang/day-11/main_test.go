package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcess(t *testing.T) {
	testCases := []struct {
		Tape          []int64
		Ref           int64
		ExpectedValue int64
	}{
		{
			[]int64{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50},
			0,
			3500,
		},
		{
			[]int64{1, 0, 0, 0, 99},
			0,
			2,
		},
		{
			[]int64{2, 3, 0, 3, 99},
			3,
			6,
		},
		{
			[]int64{2, 4, 4, 5, 99, 0},
			5,
			9801,
		},
		{
			[]int64{1, 1, 1, 4, 99, 5, 6, 0, 99},
			0,
			30,
		},
		{
			[]int64{4, 1, 4, 1, 4, 1, 99},
			0,
			4,
		},
		{
			[]int64{1002, 4, 3, 4, 33},
			4,
			99,
		},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprintf("test-%+v", c.Tape), func(tt *testing.T) {
			cpu := NewInterCode(c.Tape)
			cpu.Process()
			assert.Equal(tt, c.ExpectedValue, cpu.GetValue(c.Ref, 0))
		})
	}
}

func TestRelativeBase(t *testing.T) {
	tape := []int64{109, 19, 99}
	cpu := NewInterCode(tape)
	cpu.Process()
	assert.Equal(t, int64(19), cpu.relativeAddress)
}


func TestCopyItself(t *testing.T) {
	tape := []int64{109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99}
	cpu := NewInterCode(tape)
	cpu.Process()
}

func TestLargeNumber(t *testing.T) {
	tape := []int64{104,1125899906842624,99}
	cpu := NewInterCode(tape)
	cpu.Process()
}

func Test16digit(t *testing.T) {
	tape := []int64{1102,34915192,34915192,7,4,7,99,0}
	cpu := NewInterCode(tape)
	cpu.Process()
}