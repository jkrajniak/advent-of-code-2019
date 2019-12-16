package main

import (
	"github.com/jkrajniak/advent-of-code-2019/pkg/points"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcVelocity(t *testing.T) {
	moons := []Moon{
		{Pos: points.Point3D{3, 0, 0}},
		{Pos: points.Point3D{5, 0, 0}},
	}

	sim := NewSimulator(moons)
	sim.calcVelocity()

	assert.Equal(t, int64(1), sim.Moons[0].Vel.X)
	assert.Equal(t, int64(-1), sim.Moons[1].Vel.X)
}

func TestSimulator_Step(t *testing.T) {
	moons := []Moon{
		{Pos: points.Point3D{-1, 0, 2}},
		{Pos: points.Point3D{2, -10, -7}},
		{Pos: points.Point3D{4, -8, 8}},
		{Pos: points.Point3D{3, 5, -1}},
	}

	sim := NewSimulator(moons)
	sim.Step()

	t.Run("check velocities after one step", func(tt *testing.T) {
		assert.Equal(tt, points.Point3D{3, -1, -1}, sim.Moons[0].Vel)
		assert.Equal(tt, points.Point3D{1, 3, 3}, sim.Moons[1].Vel)
		assert.Equal(tt, points.Point3D{-3, 1, -3}, sim.Moons[2].Vel)
		assert.Equal(tt, points.Point3D{-1, -3, 1}, sim.Moons[3].Vel)
	})

	t.Run("check positions after one step", func(tt *testing.T) {
		assert.Equal(tt, points.Point3D{2, -1, 1}, sim.Moons[0].Pos)
		assert.Equal(tt, points.Point3D{3, -7, -4}, sim.Moons[1].Pos)
		assert.Equal(tt, points.Point3D{1, -7, 5}, sim.Moons[2].Pos)
		assert.Equal(tt, points.Point3D{2, 2, 0}, sim.Moons[3].Pos)
	})
}

func TestSimulator_Step2(t *testing.T) {
	moons := []Moon{
		{Pos: points.Point3D{-1, 0, 2}},
		{Pos: points.Point3D{2, -10, -7}},
		{Pos: points.Point3D{4, -8, 8}},
		{Pos: points.Point3D{3, 5, -1}},
	}

	sim := NewSimulator(moons)
	sim.Step()

	t.Run("check velocities after one step", func(tt *testing.T) {
		assert.Equal(tt, points.Point3D{3, -1, -1}, sim.Moons[0].Vel)
		assert.Equal(tt, points.Point3D{1, 3, 3}, sim.Moons[1].Vel)
		assert.Equal(tt, points.Point3D{-3, 1, -3}, sim.Moons[2].Vel)
		assert.Equal(tt, points.Point3D{-1, -3, 1}, sim.Moons[3].Vel)
	})

	t.Run("check positions after one step", func(tt *testing.T) {
		assert.Equal(tt, points.Point3D{2, -1, 1}, sim.Moons[0].Pos)
		assert.Equal(tt, points.Point3D{3, -7, -4}, sim.Moons[1].Pos)
		assert.Equal(tt, points.Point3D{1, -7, 5}, sim.Moons[2].Pos)
		assert.Equal(tt, points.Point3D{2, 2, 0}, sim.Moons[3].Pos)
	})

	// next step
	//pos=<x= 5, y=-3, z=-1>, vel=<x= 3, y=-2, z=-2>
	//	pos=<x= 1, y=-2, z= 2>, vel=<x=-2, y= 5, z= 6>
	//	pos=<x= 1, y=-4, z=-1>, vel=<x= 0, y= 3, z=-6>
	//	pos=<x= 1, y=-4, z= 2>, vel=<x=-1, y=-6, z= 2>
	sim.Step()
	t.Run("check velocities after two step", func(tt *testing.T) {
		assert.Equal(tt, points.Point3D{3, -2, -2}, sim.Moons[0].Vel)
		assert.Equal(tt, points.Point3D{-2, 5, 6}, sim.Moons[1].Vel)
		assert.Equal(tt, points.Point3D{0, 3, -6}, sim.Moons[2].Vel)
		assert.Equal(tt, points.Point3D{-1, -6, 2}, sim.Moons[3].Vel)
	})

	t.Run("check positions after two step", func(tt *testing.T) {
		assert.Equal(tt, points.Point3D{5, -3, -1}, sim.Moons[0].Pos)
		assert.Equal(tt, points.Point3D{1, -2, 2}, sim.Moons[1].Pos)
		assert.Equal(tt, points.Point3D{1, -4, -1}, sim.Moons[2].Pos)
		assert.Equal(tt, points.Point3D{1, -4, 2}, sim.Moons[3].Pos)
	})
}

func TestSimulator_Run(t *testing.T) {
	moons := []Moon{
		{Pos: points.Point3D{-1, 0, 2}},
		{Pos: points.Point3D{2, -10, -7}},
		{Pos: points.Point3D{4, -8, 8}},
		{Pos: points.Point3D{3, 5, -1}},
	}

	sim := NewSimulator(moons)

	sim.Run(10)
	energy := sim.CalculateEnergy()
	assert.Equal(t, int64(179), energy)
}