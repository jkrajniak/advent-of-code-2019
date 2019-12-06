package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"sort"
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
	for _, pA := range pointsA {
		for _, pB := range pointsB {
			if pA.Equal(pB) && !pA.Equal(Point{0, 0}){
				crossPoints = append(crossPoints, pA)
			}
		}
	}

	return crossPoints
}

func GetClosestDistanceToCentral(points []Point) (int, Point) {
	lastDistance := math.MaxInt64
	centralPoint := Point{0, 0}
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
	points := []Point{start}
	for _, cmd := range cmds {
		p := GetSegmentPoints(lastPoint, cmd)
		lastPoint = p[len(p)-1]
		points = append(points, p[1:]...)
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

func FindPathLengthToPoint(point Point, path []Point) int {
	for i := 0; i < len(path); i = i + 1 {
		if path[i].Equal(point) {
			return i
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	linesBuffer := []string{}
	idx := 0
	for scanner.Scan() {
		linesBuffer = append(linesBuffer, scanner.Text())
		if idx%2 == 1 {
			var wg sync.WaitGroup
			wg.Add(2)
			var paths [][]Point
			for _, l := range linesBuffer {
				go func(ll string) {
					defer wg.Done()
					cmd := ConvertToCommands(ll)
					p := GetPathPoints(Point{0, 0}, cmd)
					paths = append(paths, p)
				}(l)
			}
			wg.Wait()
			crossPoints := FindCrossingPoints(paths[0], paths[1])
			logrus.Debug("calculate closest distance to central")
			closestDistance, _ := GetClosestDistanceToCentral(crossPoints)
			fmt.Printf("distance: %d\n", closestDistance)

			wg.Add(len(crossPoints))
			var distancesToCrossing []int
			for _, crPoint := range crossPoints {
				go func(p Point) {
					defer wg.Done()

					numStepsA := FindPathLengthToPoint(p, paths[0])
					numStepsB := FindPathLengthToPoint(p, paths[1])
					distancesToCrossing = append(distancesToCrossing, numStepsA+numStepsB)
				}(crPoint)
			}
			wg.Wait()
			sort.Ints(distancesToCrossing)

			fmt.Printf("min num steps to intersection: %d\n", distancesToCrossing[0])

			linesBuffer = []string{}
		}
		idx = idx + 1
	}
}
