package points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoint_MulMatrix_Left(t *testing.T) {
	p := Point2D{0, 1}

	pp := p.MulMatrix(leftRotation)
	assert.Equal(t, Point2D{-1, 0}, pp)

	pp = pp.MulMatrix(leftRotation)
	assert.Equal(t, Point2D{0, -1}, pp)

	pp = pp.MulMatrix(leftRotation)
	assert.Equal(t, Point2D{1, 0}, pp)

	pp = pp.MulMatrix(leftRotation)
	assert.Equal(t, Point2D{0, 1}, pp)
}

func TestPoint_MulMatrix_Right(t *testing.T) {
	p := Point2D{0, 1}

	pp := p.MulMatrix(rightRotation)
	assert.Equal(t, Point2D{1, 0}, pp)

	pp = pp.MulMatrix(rightRotation)
	assert.Equal(t, Point2D{0, -1}, pp)

	pp = pp.MulMatrix(rightRotation)
	assert.Equal(t, Point2D{-1, 0}, pp)

	pp = pp.MulMatrix(rightRotation)
	assert.Equal(t, Point2D{0, 1}, pp)
}
