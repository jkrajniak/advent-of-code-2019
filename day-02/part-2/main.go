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
		initialMemory, err := convertToIntSlice(splittedLine)
		if err != nil {
			logrus.WithError(err).Error("failed convert to array of ints")
		}
		expectedValue := 19690720
		for n := 0; n < 100; n = n + 1 {
			for v := 0; v < 100; v = v + 1 {
				tape := append([]int{}, initialMemory...)
				tape[1] = n
				tape[2] = v
				Process(tape)
				if tape[0] == expectedValue {
					fmt.Printf("tape = %v\n", tape)
					fmt.Printf("answer = %d\n", 100*n+v)
					return
				}
			}
		}
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
