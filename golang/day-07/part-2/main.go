package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
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

type Compute struct {
	name         string
	computeGroup string
	tape         []int
	pc           int
	Input        chan int
	Output       chan int
	phase        int

	initialized bool
}

func NewCompute(name, group string, tape []int, input chan int, phase int, outOpt chan int) *Compute {
	tapeCopy := make([]int, len(tape))
	copy(tapeCopy, tape)
	var out chan int
	if outOpt != nil {
		out = outOpt
	} else {
		out = make(chan int, 10)
	}
	return &Compute{
		name,
		group,
		tapeCopy,
		0,
		input,
		out,
		phase,
		false,
	}
}

func (c *Compute) Tick() bool {
	loopCounter := 0
	for {
		opCode := NewOpCode(c.tape[c.pc])
		fmt.Println(c.name, c.computeGroup, opCode)
		switch opCode.Code {
		case 1:
			pA := c.tape[c.pc+1]
			pB := c.tape[c.pc+2]
			pC := c.tape[c.pc+3]
			c.tape[pC] = GetValue(c.tape, pA, opCode.ModeA) + GetValue(c.tape, pB, opCode.ModeB)
			c.pc += 4
		case 2:
			pA := c.tape[c.pc+1]
			pB := c.tape[c.pc+2]
			pC := c.tape[c.pc+3]
			c.tape[pC] = GetValue(c.tape, pA, opCode.ModeA) * GetValue(c.tape, pB, opCode.ModeB)
			c.pc += 4
		case 3:
			pA := c.tape[c.pc+1]
			if !c.initialized {
				c.tape[pA] = c.phase
				c.initialized = true
			} else {
				c.tape[pA] = <-c.Input
			}
			c.pc += 2
		case 4:
			pA := c.tape[c.pc+1]
			val := GetValue(c.tape, pA, opCode.ModeA)
			c.Output <- val
			c.pc += 2
		case 5:
			pA := c.tape[c.pc+1]
			pB := c.tape[c.pc+2]
			valA := GetValue(c.tape, pA, opCode.ModeA)
			valB := GetValue(c.tape, pB, opCode.ModeB)
			if valA != 0 {
				c.pc = valB
				loopCounter++
			} else {
				c.pc += 3
			}
		case 6:
			pA := c.tape[c.pc+1]
			pB := c.tape[c.pc+2]
			valA := GetValue(c.tape, pA, opCode.ModeA)
			valB := GetValue(c.tape, pB, opCode.ModeB)
			if valA == 0 {
				c.pc = valB
			} else {
				c.pc += 3
			}
		case 7:
			pA := c.tape[c.pc+1]
			pB := c.tape[c.pc+2]
			pC := c.tape[c.pc+3]
			valA := GetValue(c.tape, pA, opCode.ModeA)
			valB := GetValue(c.tape, pB, opCode.ModeB)
			if valA < valB {
				c.tape[pC] = 1
			} else {
				c.tape[pC] = 0
			}
			c.pc += 4
		case 8:
			pA := c.tape[c.pc+1]
			pB := c.tape[c.pc+2]
			pC := c.tape[c.pc+3]
			valA := GetValue(c.tape, pA, opCode.ModeA)
			valB := GetValue(c.tape, pB, opCode.ModeB)
			if valA == valB {
				c.tape[pC] = 1
			} else {
				c.tape[pC] = 0
			}
			c.pc += 4
		case 99:
			return false
		default:
			panic(fmt.Sprintf("%s: unknown opcode %d", c.name, opCode))
		}
	}
	return false
}

func ComputeThrust(tape []int, permSequence []int) int {
	inCompute := make(chan int, 10)

	//prepare 5 amplifiers, wire them together
	amp1 := NewCompute("amp1", fmt.Sprint(permSequence), tape, inCompute, permSequence[0], nil)
	amp2 := NewCompute("amp2", fmt.Sprint(permSequence), tape, amp1.Output, permSequence[1], nil)
	amp3 := NewCompute("amp3", fmt.Sprint(permSequence), tape, amp2.Output, permSequence[2], nil)
	amp4 := NewCompute("amp4", fmt.Sprint(permSequence), tape, amp3.Output, permSequence[3], nil)
	amp5 := NewCompute("amp5", fmt.Sprint(permSequence), tape, amp4.Output, permSequence[4], inCompute)

	var wg sync.WaitGroup
	wg.Add(5)

	ampStates := []int{0, 0, 0, 0, 0}
	go Runner(amp1, 0, ampStates)(&wg)
	go Runner(amp2, 1, ampStates)(&wg)
	go Runner(amp3, 2, ampStates)(&wg)
	go Runner(amp4, 3, ampStates)(&wg)
	go Runner(amp5, 4, ampStates)(&wg)

	inCompute <- 0
	wg.Wait()

	var outAmp5 int
	outAmp5 = <-amp5.Output
	return outAmp5
}

func Runner(compute *Compute, ampID int, states []int) func(group *sync.WaitGroup) {
	return func(wg *sync.WaitGroup) {
		defer func() {
			states[ampID] = 1
			wg.Done()
		}()
		compute.Tick()
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

	maxThrust := 0

	for permSequence := range getPermutations([]int{5, 6, 7, 8, 9}) {
		thrust := ComputeThrust(tape, permSequence)
		if thrust > maxThrust {
			maxThrust = thrust
		}
		//fmt.Println(permSequence)
		fmt.Println("seq", permSequence, "thrust", thrust, "maxThrust", maxThrust)
		//break
	}
}
