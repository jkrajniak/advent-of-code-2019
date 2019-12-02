package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

func GetFuel(mass int64) int64 {
	d := float64(mass) / 3.0
	return int64(math.Floor(d)) - 2
}


func GetFuelExtended(mass int64) int64 {
	fuel := GetFuel(mass)
	totalFuel := fuel
	for {
		fuel = GetFuel(fuel)
		if fuel <= 0 {
			break
		}
		totalFuel = totalFuel + fuel
	}

	return totalFuel
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	totalFuel := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			logrus.WithError(err).Errorf("failed to convert line %s to int", line)
			continue
		}
		totalFuel = totalFuel + GetFuelExtended(int64(mass))
	}
	fmt.Println(totalFuel)
}
