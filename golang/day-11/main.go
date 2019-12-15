package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	BLACK = 0
	WHITE = 1
)

type Rotation [2][2]int64

var (
	leftRotation  Rotation = [2][2]int64{{0, -1}, {1, 0}}
	rightRotation Rotation = [2][2]int64{{0, 1}, {-1, 0}}
)

func main() {
	fileName := flag.String("file", "input.txt", "input file")
	flag.Parse()

	line, err := ioutil.ReadFile(*fileName)
	if err != nil {
		panic(err)
	}

	splitLine := strings.Split(string(line), ",")
	tape, err := convertToIntSlice(splitLine)
	if err != nil {
		logrus.WithError(err).Error("failed convert to array of ints")
	}

	inCh := make(chan int64, 1)
	outCh := make(chan int64)
	cpu := NewInterCode(tape, inCh, outCh)

	direction := Point{0, -1}
	position := Point{0, 0}

	board := map[Point]int64{position: WHITE}
	visited := map[Point]int64{position: 1}

	go cpu.Process()

	for {
		inCh <- board[position]

		// read paint
		paint, ok := <-outCh
		if !ok {
			break
		}

		nextRotation, ok := <-outCh
		if !ok {
			break
		}

		// paint
		board[position] = paint
		visited[position]++
		// rotate
		if nextRotation == 0 {
			direction = direction.MulMatrix(leftRotation)
		} else if nextRotation == 1 {
			direction = direction.MulMatrix(rightRotation)
		}
		position = position.Add(direction)

	}

	fmt.Printf("num of painted tails = %d", len(board))

	minX, maxX := 0, 0
	minY, maxY := 0, 0
	for p, _ := range board {
		minX = int(math.Min(float64(minX), float64(p.X)))
		maxX = int(math.Max(float64(maxX), float64(p.X)))
		minY = int(math.Min(float64(minY), float64(p.Y)))
		maxY = int(math.Max(float64(maxY), float64(p.Y)))
	}

	fmt.Printf("X = %d:%d, Y = %d:%d", minX, maxX, minY, maxY)

	fmt.Println("\nBoard:")
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if v, ok := board[Point{int64(x), int64(y)}]; ok {
				c := '#'
				if v == 0 {
					c = ' '
				}
				fmt.Printf("%c", c)
			}
		}
		fmt.Println()
	}
}

func convertToIntSlice(s []string) ([]int64, error) {
	var ints []int64
	for _, s := range s {
		i, err := strconv.Atoi(strings.Trim(s, "\n"))
		if err != nil {
			return nil, err
		}
		ints = append(ints, int64(i))
	}
	return ints, nil
}
