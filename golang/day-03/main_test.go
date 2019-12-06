package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersections(t *testing.T) {
	testCases := []struct {
		Paths    []string
		Distance int
	}{
		{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
		{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135},
	}
	for i, c := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			cmdsA := ConvertToCommands(c.Paths[0])
			cmdsB := ConvertToCommands(c.Paths[1])
			pathA := GetPathPoints(Point{0, 0}, cmdsA)
			pathB := GetPathPoints(Point{0, 0}, cmdsB)
			crossPoints := FindCrossingPoints(pathA, pathB)
			//assert.Equal(t, "", crossPoints)
			closestDistance, _ := GetClosestDistanceToCentral(crossPoints)
			assert.Equal(tt, c.Distance, closestDistance)
		})
	}
}

func TestCrossingPoints(t *testing.T) {
	pointsA := GetSegmentPoints(Point{0, 0}, Command{
		Direction: "R",
		Steps:     5,
	})
	assert.Equal(t, 6, len(pointsA))
	assert.Equal(t, []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}}, pointsA)
	pointsB := GetSegmentPoints(Point{1, 5}, Command{
		Direction: "D",
		Steps:     8,
	})
	assert.Equal(t, 9, len(pointsB))
	crossPoints := FindCrossingPoints(pointsA, pointsB)
	assert.Equal(t, []Point{{1, 0}}, crossPoints)
}

func TestCrossingPoints_ShortA(t *testing.T) {
	pointsA := GetSegmentPoints(Point{0, 0}, Command{
		Direction: "R",
		Steps:     2,
	})
	assert.Equal(t, 3, len(pointsA))
	pointsB := GetSegmentPoints(Point{1, 1}, Command{
		Direction: "D",
		Steps:     8,
	})
	assert.Equal(t, 9, len(pointsB))
	crossPoints := FindCrossingPoints(pointsA, pointsB)
	assert.Equal(t, []Point{{1, 0}}, crossPoints)
}

func TestNumStepsToReachPoint(t *testing.T) {
	testCases := []struct {
		RAWCmd   string
		NumSteps int
		Intersect Point
	}{
		{"R8,U5,L5,D3", 20, Point{3, 3}},
		{"U7,R6,D4,L4", 20, Point{3, 3}},
		{"R8,U5,L5,D3", 15, Point{6, 5}},
		{"U7,R6,D4,L4", 15, Point{6, 5}},
	}
	for i, c := range testCases {
		t.Run(fmt.Sprintf("test-%d", i), func(tt *testing.T) {
			cmds := ConvertToCommands(c.RAWCmd)
			path := GetPathPoints(Point{}, cmds)
			numSteps := FindPathLengthToPoint(c.Intersect, path)
			assert.Equal(tt, c.NumSteps, numSteps)
		})
	}
}
