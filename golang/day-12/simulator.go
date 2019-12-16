package main

type Simulator struct {
	Moons []Moon
	currentStep int64
}

func NewSimulator(moons []Moon) *Simulator {
	sim := &Simulator{Moons: moons}
	return sim
}


func (m *Simulator) calcVelocity() {
	for i := 0; i < len(m.Moons); i++ {
		for j := i + 1; j < len(m.Moons); j++ {
			mooni := &m.Moons[i]
			moonj := &m.Moons[j]

			for ax := 0; ax < 3; ax++ {
				if mooni.Pos.At(ax) == moonj.Pos.At(ax) {
					continue
				}

				if mooni.Pos.At(ax) < moonj.Pos.At(ax) {
					mooni.Vel.Set(ax, mooni.Vel.At(ax) + 1)
					moonj.Vel.Set(ax, moonj.Vel.At(ax) - 1)
				} else {
					mooni.Vel.Set(ax, mooni.Vel.At(ax) - 1)
					moonj.Vel.Set(ax, moonj.Vel.At(ax) + 1)
				}
			}
		}
	}
}

func (m *Simulator) calcPosition() {
	for i, _ := range m.Moons {
		m := &m.Moons[i]
		m.Pos.Add(m.Vel)
	}
}

func (m *Simulator) Step() {
	m.calcVelocity()
	m.calcPosition()
	m.currentStep++
}

func (m *Simulator) Run(n int64) {
	for m.currentStep < n {
		m.Step()
	}
}

func (m *Simulator) CalculateEnergy() int64 {
	totalEnergy := int64(0)
	for _, m := range m.Moons {
		totalEnergy = totalEnergy + (abs(m.Pos.X) + abs(m.Pos.Y) + abs(m.Pos.Z)) * (abs(m.Vel.X) + abs(m.Vel.Y) + abs(m.Vel.Z))
	}
	return totalEnergy
}

func abs(x int64) int64 {
	if x >= 0 {
		return x
	}
	return -1*x
}