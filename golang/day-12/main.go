package main

import (
	"bufio"
	"fmt"
	"github.com/jkrajniak/advent-of-code-2019/pkg/points"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	moons := Read(scanner)

	initialState := make([]Moon, len(moons))
	copy(initialState, moons)

	initialMoonMap := map[int64]Moon{}
	for _, m := range initialState {
		initialMoonMap[m.ID] = m
	}

	simulator := NewSimulator(moons)

	// periods
	periods := points.Point3D{0, 0, 0}

	// loop
	for s := int64(1); periods.X == 0 || periods.Y == 0 || periods.Z == 0; s++ {
		simulator.Step()
		validMoons := [3]int{}
		for _, m := range simulator.Moons {
			initMoon := initialMoonMap[m.ID]
			for ax := 0; ax < 3; ax++ {
				if periods.At(ax) == 0 && initMoon.Pos.At(ax) == m.Pos.At(ax) && initMoon.Vel.At(ax) == m.Vel.At(ax) {
					validMoons[ax]++
				}
			}
		}
		for ax := 0; ax < 3; ax++ {
			if validMoons[ax] == len(moons) {
				periods.Set(ax, s)
				fmt.Println(s, periods, validMoons)
			}
		}
	}
	fmt.Println(periods)
	fmt.Println("lcm = ", LCM(periods.X, periods.Y, periods.Z))
}

func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int64, integers ...int64) int64 {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
