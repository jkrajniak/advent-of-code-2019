package points

type Point3D struct {
	X, Y, Z int64
}

func (p *Point3D) Add(b Point3D) {
	p.X += b.X
	p.Y += b.Y
	p.Z += b.Z
}

func (p *Point3D) At(idx int) int64 {
	switch idx {
	case 0:
		return p.X
	case 1:
		return p.Y
	case 2:
		return p.Z
	default:
		panic("wrong index")
	}
}

func (p *Point3D) Set(idx int, val int64) {
	switch idx {
	case 0:
		p.X = val
	case 1:
		p.Y = val
	case 2:
		p.Z = val
	default:
		panic("wrong index")
	}
}
