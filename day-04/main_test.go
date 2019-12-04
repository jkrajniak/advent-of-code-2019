package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPassword(t *testing.T) {
	testCase := []struct{
		PassCode []int
		Expected bool
	}{
		{[]int{1,1,1,1,1,1}, true},
		{[]int{2,2,3,4,5,0}, false},
		{[]int{1,2,3,7,8,9}, false},
		{[]int{1}, false},
		{[]int{1,1,1,1,2,3}, true},
		{[]int{1,3,5,6,7,9}, false},
	}
	for i, c := range testCase {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			assert.Equal(tt, c.Expected, CheckPassword(c.PassCode))
		})
	}
}

func TestCheckPasswordExtended(t *testing.T) {
	testCase := []struct{
		PassCode []int
		Expected bool
	}{
		{[]int{1,1,2,2,3,3}, true},
		{[]int{1,2,3,4,4,4}, false},
		{[]int{1,1,1,1,2,2}, true},
		{[]int{1,1,1,1,1,2}, false},
		{[]int{1,1,1,1,1,1}, false},
		{[]int{4,4,4,5,6,7}, false},
	}
	for _, c := range testCase {
		t.Run(fmt.Sprintf("passcode-%v", c.PassCode), func(tt *testing.T) {
			assert.Equal(tt, c.Expected, CheckPasswordExtended(c.PassCode))
		})
	}
}

func TestConvertToArrayOfInts(t *testing.T) {
	testCase := []struct{
		Code int64
		Expected []int
	}{
		{1111, []int{1,1,1,1}},
		{1, []int{1}},
		{123456, []int{1,2,3,4,5,6}},
	}
	for _, c := range testCase {
		t.Run(fmt.Sprintf("test-%d", c.Code), func(tt *testing.T) {
			code := convertToArrayOfInts(c.Code)
			assert.Equal(t, fmt.Sprint(c.Expected), fmt.Sprint(code))
		})
	}
}

func TestGetPassCodes(t *testing.T) {
	testCase := []struct{
		Start, End int
		NumPassCodes int
	}{
		{284639, 748759, 895},
		{100000, 111111, 1},
		{100000, 111112, 2},
		{100000, 121314, 495},
	}

	for _, c := range testCase {
		t.Run(fmt.Sprintf("test-%d:%d", c.Start, c.End), func(tt *testing.T) {
			passCodes := GetPassCodes(c.Start, c.End)
			assert.Equal(tt, c.NumPassCodes, len(passCodes))
		})
	}
}