package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
)

type OpCode struct {
	Code  int
	ModeA int
	ModeB int
	ModeC int
}

func NewOpCode(val int) OpCode {
	s := fmt.Sprintf("%05d", val)
	opCode, _ := strconv.Atoi(s[3:])
	modeA, _ := strconv.Atoi(s[2:3])
	modeB, _ := strconv.Atoi(s[1:2])
	modeC, _ := strconv.Atoi(s[0:1])
	return OpCode{
		Code:  opCode,
		ModeA: modeA,
		ModeB: modeB,
		ModeC: modeC,
	}
}

func GetValue(tape []int, pos, mode int) int {
	if mode == 1 {
		return pos
	}
	return tape[pos]
}

func Process(tape []int) {
	pc := 0
	for pc < len(tape) {
		opCode := NewOpCode(tape[pc])
		//fmt.Println(opCode)
		switch opCode.Code {
		case 1:
			pA := tape[pc+1]
			pB := tape[pc+2]
			pC := tape[pc+3]
			tape[pC] = GetValue(tape, pA, opCode.ModeA) + GetValue(tape, pB, opCode.ModeB)
			pc += 4
		case 2:
			pA := tape[pc+1]
			pB := tape[pc+2]
			pC := tape[pc+3]
			tape[pC] = GetValue(tape, pA, opCode.ModeA) * GetValue(tape, pB, opCode.ModeB)
			pc += 4
		case 3:
			pA := tape[pc+1]
			var in int
			fmt.Scanln(&in)
			tape[pA] = in
			pc += 2
		case 4:
			pA := tape[pc+1]
			val := GetValue(tape, pA, opCode.ModeA)
			fmt.Println(val)
			pc += 2
		case 5:
			pA := tape[pc+1]
			pB := tape[pc+2]
			valA := GetValue(tape, pA, opCode.ModeA)
			valB := GetValue(tape, pB, opCode.ModeB)
			if valA != 0 {
				pc = valB
			} else {
				pc += 3
			}
		case 6:
			pA := tape[pc+1]
			pB := tape[pc+2]
			valA := GetValue(tape, pA, opCode.ModeA)
			valB := GetValue(tape, pB, opCode.ModeB)
			if valA == 0 {
				pc = valB
			} else {
				pc += 3
			}
		case 7:
			pA := tape[pc+1]
			pB := tape[pc+2]
			pC := tape[pc+3]
			valA := GetValue(tape, pA, opCode.ModeA)
			valB := GetValue(tape, pB, opCode.ModeB)
			if valA < valB {
				tape[pC] = 1
			} else {
				tape[pC] = 0
			}
			pc += 4
		case 8:
			pA := tape[pc+1]
			pB := tape[pc+2]
			pC := tape[pc+3]
			valA := GetValue(tape, pA, opCode.ModeA)
			valB := GetValue(tape, pB, opCode.ModeB)
			if valA == valB {
				tape[pC] = 1
			} else {
				tape[pC] = 0
			}
			pc += 4
		case 99:
			return
		}
	}
}

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
	Process(tape)
}

func convertToIntSlice(s []string) ([]int, error) {
	var ints []int
	for _, s := range s {
		i, err := strconv.Atoi(strings.Trim(s, "\n"))
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
