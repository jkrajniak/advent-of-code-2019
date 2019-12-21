package main

import (
	"bufio"
	"github.com/jkrajniak/advent-of-code-2019/pkg/points"
	"regexp"
	"strconv"
)

var (
	moonsPositionExp = regexp.MustCompile(`x=(-?[0-9]\d*), y=(-?[0-9]\d*), z=(-?[0-9]\d*)`)
)

type Moon struct {
	ID  int64
	Pos points.Point3D
	Vel points.Point3D
}

func mustInt(s string) int64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return int64(i)
}

func Read(scanner *bufio.Scanner) []Moon {
	var moons []Moon
	moonID := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		result := moonsPositionExp.FindStringSubmatch(line)
		moons = append(moons, Moon{Pos: points.Point3D{
			X: mustInt(result[1]),
			Y: mustInt(result[2]),
			Z: mustInt(result[3]),
		}, ID: moonID})
		moonID++
	}

	return moons
}
