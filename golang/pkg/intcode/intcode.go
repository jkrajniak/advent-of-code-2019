package intcode

import (
	"fmt"
	"strconv"
)

type OpCode struct {
	Code  int
	ModeA int
	ModeB int
	ModeC int
}

func NewOpCode(val int64) OpCode {
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

type Tape struct {
	tape []int64
}

func NewTape(tape []int64) *Tape {
	copyTape := make([]int64, len(tape))
	copy(copyTape, tape)
	return &Tape{copyTape}
}

func (t *Tape) Write(val, pos int64) {
	if pos >= int64(len(t.tape)) {
		// slice is full, increase the capacity by a factor of 2
		t.extendTape(pos)
	}

	t.tape[pos] = val
}

func (t *Tape) extendTape(pos int64) {
	newSlice := make([]int64, pos+1, 2*pos+1)
	copy(newSlice, t.tape)
	t.tape = newSlice
}

func (t *Tape) Read(pos int64) int64 {
	if pos >= int64(cap(t.tape)) {
		t.extendTape(pos)
	}
	return t.tape[pos]
}

func (t *Tape) Len() int64 { return int64(len(t.tape)) }

type InterCode struct {
	pc              int64
	relativeAddress int64
	tape            *Tape

	inCh  <-chan int64
	outCh chan<- int64
	promptInputCh chan<- bool
}

func NewInterCode(tape []int64, inCh <-chan int64, outCh chan<- int64, promptInput chan<- bool) *InterCode {
	return &InterCode{tape: NewTape(tape), inCh: inCh, outCh: outCh, promptInputCh:promptInput}
}

func (i *InterCode) GetValue(pos int64, mode int) int64 {
	switch mode {
	case 1:
		return pos
	case 2:
		return i.tape.Read(pos + i.relativeAddress)
	default:
		return i.tape.Read(pos)
	}
}

func (i *InterCode) WriteValue(val, pos int64, mode int) {
	writePos := pos
	if mode == 2 {
		writePos = i.relativeAddress + pos
	}
	i.tape.Write(val, writePos)
}

func (i *InterCode) Next() bool {
	opCode := NewOpCode(i.tape.Read(i.pc))
	//fmt.Println(opCode)
	switch opCode.Code {
	case 1:
		pA := i.tape.Read(i.pc + 1)
		pB := i.tape.Read(i.pc + 2)
		pC := i.tape.Read(i.pc + 3)
		valA := i.GetValue(pA, opCode.ModeA)
		valB := i.GetValue(pB, opCode.ModeB)
		i.WriteValue(valA+valB, pC, opCode.ModeC)
		i.pc += 4
	case 2:
		pA := i.tape.Read(i.pc + 1)
		pB := i.tape.Read(i.pc + 2)
		pC := i.tape.Read(i.pc + 3)
		valA := i.GetValue(pA, opCode.ModeA)
		valB := i.GetValue(pB, opCode.ModeB)
		i.WriteValue(valA*valB, pC, opCode.ModeC)
		i.pc += 4
	case 3:
		pA := i.tape.Read(i.pc + 1)
		if i.promptInputCh != nil {
			i.promptInputCh <- true
		}
		i.WriteValue(<-i.inCh, pA, opCode.ModeA)
		i.pc += 2
	case 4:
		pA := i.tape.Read(i.pc + 1)
		val := i.GetValue(pA, opCode.ModeA)
		i.outCh <- val
		i.pc += 2
	case 5:
		pA := i.tape.Read(i.pc + 1)
		pB := i.tape.Read(i.pc + 2)
		valA := i.GetValue(pA, opCode.ModeA)
		valB := i.GetValue(pB, opCode.ModeB)
		if valA != 0 {
			i.pc = valB
		} else {
			i.pc += 3
		}
	case 6:
		pA := i.tape.Read(i.pc + 1)
		pB := i.tape.Read(i.pc + 2)
		valA := i.GetValue(pA, opCode.ModeA)
		valB := i.GetValue(pB, opCode.ModeB)
		if valA == 0 {
			i.pc = valB
		} else {
			i.pc += 3
		}
	case 7:
		pA := i.tape.Read(i.pc + 1)
		pB := i.tape.Read(i.pc + 2)
		pC := i.tape.Read(i.pc + 3)
		valA := i.GetValue(pA, opCode.ModeA)
		valB := i.GetValue(pB, opCode.ModeB)
		valC := 0
		if valA < valB {
			valC = 1
		}
		i.WriteValue(int64(valC), pC, opCode.ModeC)
		i.pc += 4
	case 8:
		pA := i.tape.Read(i.pc + 1)
		pB := i.tape.Read(i.pc + 2)
		pC := i.tape.Read(i.pc + 3)
		valA := i.GetValue(pA, opCode.ModeA)
		valB := i.GetValue(pB, opCode.ModeB)
		valC := 0
		if valA == valB {
			valC = 1
		}
		i.WriteValue(int64(valC), pC, opCode.ModeC)
		i.pc += 4
	case 9:
		pA := i.tape.Read(i.pc + 1)
		valA := i.GetValue(pA, opCode.ModeA)
		i.relativeAddress += valA
		i.pc += 2
	case 99:
		return true
	}
	return false
}

func (i *InterCode) Process() bool {
	i.pc = 0
	for i.pc < i.tape.Len() {
		exit := i.Next()
		if exit {
			if i.outCh != nil {
				close(i.outCh)
			}
			return true
		}
	}
	return false
}
