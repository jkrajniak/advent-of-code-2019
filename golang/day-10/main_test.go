package main

import (
	"bufio"
	"fmt"
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

func TestIsPointOnLine(t *testing.T) {
	testCases := []struct {
		A        Point
		B        Point
		P        Point
		Expected bool
	}{{
		Point{0, 0},
		Point{10, 0},
		Point{5, 0},
		true,
	}, {
		Point{0, 0},
		Point{10, 0},
		Point{5, 5},
		false,
	},
	}

	for _, c := range testCases {
		t.Run(fmt.Sprintf("test-%+v:%+v", c.A, c.B), func(tt *testing.T) {
			assert.Equal(tt, c.Expected, IsPointOnLine(c.A, c.B, c.P))
		})
	}
}

func TestCalculateAsteroids(t *testing.T) {
	testCases := []struct {
		Map          string
		NumAsteroids int64
	}{
		{`.#..#
.....
#####
....#
...##`, 8},
		{
			`......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`, 33,
		},
		{`.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`, 210},
	}
	for i, c := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(c.Map))
			points := convertToListOfPoints(scanner)
			num, p := GetCountOfMaxObservableAsteroids(points)
			assert.Equal(tt, c.NumAsteroids, num)
			assert.NotNil(tt, p)
		})
	}

}

func TestAngle(t *testing.T) {
	p := []Point{
		{-1, 0},
		{2, 2},
		{1, 1},
		{1, 0},
		{0, -1},
		{0, 1},
		{-1, -1},
		{1, -1},
		{-1, 1}}

	expected := []Point{
		{0, 1},
		{1, 1},
		{2, 2},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
	}

	fmt.Println(p)
	pp2 := SortPoints(p, Point{0, 0})
	fmt.Println("sortPoints", pp2)

	assert.Equal(t, expected, pp2)
}
