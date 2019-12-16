package main

import (
	"bufio"
	"github.com/jkrajniak/advent-of-code-2019/pkg/points"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	s := `<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>`
	scanner := bufio.NewScanner(strings.NewReader(s))
	moons := Read(scanner)

	assert.Equal(t, points.Point3D{-1, 0, 2}, moons[0].Pos)
	assert.Equal(t, points.Point3D{2, -10, -7}, moons[1].Pos)
	assert.Equal(t, points.Point3D{4, -8, 8}, moons[2].Pos)
	assert.Equal(t, points.Point3D{3, 5, -1}, moons[3].Pos)
}