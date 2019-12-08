package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeros(t *testing.T) {
	answer := Calculate("123456789012", 6)
	assert.Equal(t, 1, answer)
}


func TestImage(t *testing.T) {
	image := ComputeImage("0222112222120000", 4)
	assert.Equal(t, []int{0,1,1,0}, image)
}

func TestImage2(t *testing.T) {
	image := ComputeImage("02221122221222220001", 4)
	assert.Equal(t, []int{0,1,1,1}, image)
}