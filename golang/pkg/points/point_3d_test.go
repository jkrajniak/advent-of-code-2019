package points

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoint3D_Add(t *testing.T) {
	p := &Point3D{1, 2, 3}
	p.Add(Point3D{1, 2, 3})
	assert.Equal(t, int64(2), p.X)
	assert.Equal(t, int64(4), p.Y)
	assert.Equal(t, int64(6), p.Z)
}
