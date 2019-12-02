package main

import (
	"bufio"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

func Process(tape []int) {
	for i := 0; i < len(tape)-4; i = i + 4 {
		opCode := tape[i]
		pA := tape[i+1]
		pB := tape[i+2]
		pC := tape[i+3]
		switch opCode {
		case 1:
			tape[pC] = tape[pA] + tape[pB]
		case 2:
			tape[pC] = tape[pA] * tape[pB]
		case 99:
			return
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		splittedLine := strings.Split(scanner.Text(), ",")
		tape, err := convertToIntSlice(splittedLine)
		if err != nil {
			logrus.WithError(err).Error("failed convert to array of ints")
		}
		Process(tape)
		fmt.Println(tape)
	}
}

func convertToIntSlice(s []string) ([]int, error) {
	var ints []int
	for _, s := range s {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}