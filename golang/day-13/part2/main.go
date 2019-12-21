package main

import (
	"flag"
	"fmt"
	"github.com/jkrajniak/advent-of-code-2019/pkg/intcode"
	"github.com/jkrajniak/advent-of-code-2019/pkg/points"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	BallID   = 4
	PaddleID = 3
)

var (
	JoystickNeutral int64 = 0
	JoystickLeft    int64 = -1
	JoystickRight   int64 = 1
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

	inCh := make(chan int64)
	outCh := make(chan int64)
	inRequest := make(chan bool)

	tape[0] = 2
	cpu := intcode.NewInterCode(tape, inCh, outCh, inRequest)
	go cpu.Process()

	outputOk := false
	x, y := int64(0), int64(0)
	inputIdx := 0
	ballPos := points.Point2D{}
	paddlePos := points.Point2D{}

	for {
		select {
		//read from channel
		case v, ok := <-outCh:
			if !ok {
				outputOk = true
				break
			}
			switch inputIdx {
			case 0:
				x = v
				inputIdx++
			case 1:
				y = v
				inputIdx++
			case 2:
				if x == -1 && y == 0 {
					fmt.Println("score", v)
				} else {
					pos := points.Point2D{x, y}
					if v == PaddleID {
						paddlePos = pos
					} else if v == BallID {
						ballPos = pos
					}
				}
				inputIdx = 0
			}
		case _, ok := <-inRequest:
			if !ok {

			}
			if ballPos.X < paddlePos.X {
				inCh <- JoystickLeft
			} else if ballPos.X > paddlePos.X {
				inCh <- JoystickRight
			} else {
				inCh <- JoystickNeutral
			}
		}
		if outputOk {
			break
		}
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
