package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

var OrientationMap = map[string]Point{
	"R": {1, 0},
	"D": {0, -1},
	"L": {-1, 0},
	"U": {0, 1},
}

type Command struct {
	Direction string
	Steps     int
}

func (c Command) GetPoint() Point {
	return OrientationMap[c.Direction].Mul(c.Steps)
}

type Point struct {
	X, Y int
}

func (p Point) Add(b Point) Point {
	return Point{p.X + b.X, p.Y + b.Y}
}

func (p Point) Mul(c int) Point {
	return Point{c * p.X, c * p.Y}
}

func (p Point) Equal(b Point) bool {
	return p.X == b.X && p.Y == b.Y
}

func GetSegmentPoints(start Point, cmd Command) []Point {
	points := []Point{start}
	lastPoint := start
	orientationPoint := OrientationMap[cmd.Direction]
	for c := 0; c < cmd.Steps; c++ {
		lastPoint = lastPoint.Add(orientationPoint)
		points = append(points, lastPoint)
	}
	return points
}

func FindCrossingPoints(pointsA []Point, pointsB []Point) []Point {
	var crossPoints []Point
	setA := pointsA
	setB := pointsB

	for a := 0; a < len(pointsA); a = a + 1 {
		for b := a + 1; b < len(pointsB); b = b + 1 {
			pA := pointsA[a]
			pB := pointsB[b]
			if pA.Equal(pB) {
				crossPoints = append(crossPoints, pA)
			}
		}
	}
	return crossPoints
}

func GetClosestDistanceToCentral(points []Point) (int, Point) {
	lastDistance := math.MaxInt64
	centralPoint := Point{0,0}
	lastPoint := centralPoint
	for _, p := range points {
		if p.Equal(centralPoint) {
			continue
		}
		distance := int(math.Abs(float64(p.X)) + math.Abs(float64(p.Y)))
		if distance < lastDistance {
			lastDistance = distance
			lastPoint = p
		}
	}
	return lastDistance, lastPoint
}

func GetPathPoints(start Point, cmds []Command) []Point {
	lastPoint := start
	points := []Point{}
	for _, cmd := range cmds {
		p := GetSegmentPoints(lastPoint, cmd)
		lastPoint = p[len(p)-1]
		points = append(points, p...)
	}
	return points
}

func ConvertToCommands(line string) []Command {
	s := strings.Split(line, ",")
	var commands []Command
	for _, c := range s {
		steps, err := strconv.Atoi(c[1:])
		if err != nil {
			panic(err)
		}
		commands = append(commands, Command{string(c[0]), steps})
	}
	return commands
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	linesBuffer := []string{}
	idx := 0
	for scanner.Scan() {
		linesBuffer = append(linesBuffer, scanner.Text())
		if idx % 2 == 1 {
			var wg sync.WaitGroup
			wg.Add(2)
			paths := [][]Point{}
			for _, l := range linesBuffer {
				go func(ll string) {
					defer wg.Done()
					cmd := ConvertToCommands(ll)
					p := GetPathPoints(Point{0, 0}, cmd)
					paths = append(paths, p)
				}(l)
			}
			wg.Wait()
			fmt.Println("ended processing, find cross points")
			crossPoints := FindCrossingPoints(paths[0], paths[1])
			closestDistance, _ := GetClosestDistanceToCentral(crossPoints)
			fmt.Printf("distance: %d\n", closestDistance)
			linesBuffer = []string{}
		}
		idx = idx + 1
	}
}
