package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

func TestGetPoints(t *testing.T) {
	s := `.#..#
.....
..#..
.....
...#.`

	scanner := bufio.NewScanner(strings.NewReader(s))
	points := convertToListOfPoints(scanner)
	expected := []Point{
		{1, 0}, {4, 0},
		{2, 2},
		{3, 4},
	}
	sort.Slice(expected, func(i, j int) bool {
		return expected[i].X < expected[j].X
	})
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X
	})
	assert.Equal(t, expected, points)
}

func TestCalculateAsteroids(t *testing.T) {
	s := `.#..#
.....
#####
....#
...##`
	scanner := bufio.NewScanner(strings.NewReader(s))
	points := convertToListOfPoints(scanner)
	num, p := GetCountOfMaxObservableAsteroids(points)
	assert.Equal(t, 8, num)
	assert.NotNil(t, p)
}
