package points

type Point2D struct {
	X, Y int64
}

func (p Point2D) MulMatrix(m [2][2]int64) Point2D {
	out := Point2D{}
	out.X = m[0][0]*p.X + m[0][1]*p.Y
	out.Y = m[1][0]*p.X + m[1][1]*p.Y
	return out
}

func (p Point2D) Add(b Point2D) Point2D {
	return Point2D{p.X + b.X, p.Y + b.Y}
}
