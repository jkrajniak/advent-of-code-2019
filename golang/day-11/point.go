package main


type Point struct {
	X, Y int64
}

func (p Point) MulMatrix(m [2][2]int64) Point {
	out := Point{}
	out.X = m[0][0]*p.X + m[0][1]*p.Y
	out.Y = m[1][0]*p.X + m[1][1]*p.Y
	return out
}

func (p Point) Add(b Point) Point {
	return Point{p.X + b.X, p.Y + b.Y}
}
