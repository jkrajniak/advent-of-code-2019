package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

func (p Point) Eq(b Point) bool {
	return p.X == b.X && p.Y == b.Y
}

func (p Point) Sub(b Point) Point {
	return Point{p.X - b.X, p.Y - b.Y}
}

func (p Point) Add(b Point) Point {
	return Point{p.X + b.X, p.Y + b.Y}
}

func (p Point) Mul(c int) Point {
	return Point{c * p.X, c * p.Y}
}

func IsPointOnLine(a, b, p Point) bool {
	// check if three points are aligned
	cross := (p.Y-a.Y)*(b.X-a.X) - (p.X-a.X)*(b.Y-a.Y)
	if cross != 0 {
		return false
	}

	dot := (p.X-a.X)*(b.X-a.X) + (p.Y-a.Y)*(b.Y-a.Y)
	if dot < 0 {
		return false
	}

	sqBA := (b.X-a.X)*(b.X-a.X) + (b.Y-a.Y)*(b.Y-a.Y)
	if dot > sqBA {
		return false
	}

	return true
}

func Angle(a, b Point) float64 {
	dot := float64(a.X*b.X + a.Y*b.Y)
	lA := math.Sqrt(float64(a.X*a.X + a.Y*a.Y))
	lB := math.Sqrt(float64(b.X*b.X + b.Y*b.Y))

	return dot / (lA * lB)
}

func SortPoints(ps []Point, com Point) []Point {
	q := make([][]Point, 4)
	for _, p := range ps {
		if p.X >= 0 && p.Y >= 0 {
			q[0] = append(q[0], p)
		} else if p.X >= 0 && p.Y < 0 {
			q[1] = append(q[1], p)
		} else if p.X < 0 && p.Y < 0 {
			q[2] = append(q[2], p)
		} else {
			q[3] = append(q[3], p)
		}
	}
	var qq []Point
	axes := []Point{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for ai, v := range q {
		sort.Slice(v, func(i, j int) bool {
			vi, vj := v[i], v[j]
			iCOMsq := (com.X-vi.X)*(com.X-vi.X) + (com.Y-vi.Y)*(com.Y-vi.Y)
			jCOMsq := (com.X-vj.X)*(com.X-vj.X) + (com.Y-vj.Y)*(com.Y-vj.Y)
			angleI := Angle(axes[ai], vi)
			angleJ := Angle(axes[ai], vj)
			if angleI == angleJ {
				return iCOMsq < jCOMsq
			}
			return angleI > angleJ
		})
		qq = append(qq, v...)
	}
	return qq
}

func MovePoints(points []Point, com Point) []Point {
	var out []Point
	for _, p := range points {
		pMoved := p.Sub(com)
		out = append(out, pMoved)
	}
	return out
}

func GetVisibleAsteroids(src Point, points []Point) []Point {
	var visiblePoints []Point
	for _, dstPoint := range points {
		if src.Eq(dstPoint) {
			continue
		}

		validDstPoint := true
		for ip := 0; ip < len(points) && validDstPoint; ip++ {
			p := points[ip]
			if p.Eq(src) || p.Eq(dstPoint) {
				continue
			}
			validDstPoint = validDstPoint && !IsPointOnLine(src, dstPoint, p)
		}
		if validDstPoint {
			visiblePoints = append(visiblePoints, dstPoint)
		}
	}
	return visiblePoints
}

func GetCountOfMaxObservableAsteroids(points []Point) (int64, *Point) {
	seenAsteroids := map[Point]int64{}

	for _, srcPoint := range points {
		for _, dstPoint := range points {
			if srcPoint.Eq(dstPoint) {
				continue
			}
			validDstPoint := true
			for ip := 0; ip < len(points) && validDstPoint; ip++ {
				p := points[ip]
				if p.Eq(srcPoint) || p.Eq(dstPoint) {
					continue
				}
				validDstPoint = validDstPoint && !IsPointOnLine(srcPoint, dstPoint, p)
			}
			if validDstPoint {
				seenAsteroids[srcPoint]++
			}
		}
	}

	maxCount := int64(0)
	var maxPoint Point

	for p, c := range seenAsteroids {
		if c > maxCount {
			maxCount = c
			maxPoint = p
		}
	}

	return maxCount, &maxPoint
}

func convertToListOfPoints(scanner *bufio.Scanner) []Point {
	var points []Point
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, v := range line {
			if v == '#' {
				points = append(points, Point{x, y})
			}
		}
		y++
	}
	return points
}

func getAvailablePoints(vaporizedPoints map[Point]bool) (out []Point) {
	for k, v := range vaporizedPoints {
		if !v {
			out = append(out, k)
		}
	}
	return out
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	points := convertToListOfPoints(scanner)
	num, com := GetCountOfMaxObservableAsteroids(points)
	fmt.Printf("num = %d, point = %+v\n", num, com)

	movedPoints := MovePoints(points, *com)
	fmt.Println(len(movedPoints))

	vaporizedPoints := map[Point]bool{}
	for _, p := range movedPoints {
		vaporizedPoints[p] = false
	}

	cidx := 1

	for cidx <= 200 {
		visibleAsteroids := GetVisibleAsteroids(Point{0, 0}, getAvailablePoints(vaporizedPoints))
		sortedVisibleAsteroids := SortPoints(visibleAsteroids, Point{0, 0})
		fmt.Println(len(visibleAsteroids))
		for _, p := range sortedVisibleAsteroids {
			vaporizedPoints[p] = true
			fmt.Printf("%d %+v\n", cidx, p.Add(*com))
			cidx = cidx + 1
			if cidx > 201 {
				break
			}
		}
	}
}
