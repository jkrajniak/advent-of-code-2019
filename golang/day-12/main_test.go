package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLCM2(t *testing.T) {
	r := LCM(2, 3, 4)
	assert.Equal(t, int64(12), r)
}