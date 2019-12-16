package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	moons := Read(scanner)

	simulator := NewSimulator(moons)

	simulator.Run(1000)
	fmt.Println(simulator.currentStep, simulator.CalculateEnergy())
}
