package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T) {
	testCases := []struct{
		Mass int64
		ExpectedFuel int64
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}
	for _, c := range testCases {
		t.Run(fmt.Sprintf("mass-%d", c.Mass), func(t *testing.T) {
			fuel := GetFuel(c.Mass)
			assert.Equal(t, c.ExpectedFuel, fuel)
		})
	}
}
