package framework

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRadian_LefterThan(t *testing.T) {
	a1 := Radian(0)
	t.Run("check left angle about angle", func(t *testing.T) {
		assert.False(t, a1.LefterThan(Degrees(-181).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(-180).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-179).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-90).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-1).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(0).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(1).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(10).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(45).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(90).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(135).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(179).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(180).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(181).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(235).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(275).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(300).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(330).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(359).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(360).ToRadians()))
	})
	a1 = Degrees(-1).ToRadians()
	t.Run("check left angle about angle", func(t *testing.T) {
		assert.False(t, a1.LefterThan(Degrees(-181).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-180).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-179).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-90).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(-1).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(0).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(1).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(10).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(45).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(90).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(135).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(179).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(180).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(181).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(235).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(275).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(300).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(330).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(359).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(360).ToRadians()))
	})

	a1 = Degrees(1).ToRadians()
	t.Run("check left angle about angle", func(t *testing.T) {
		assert.False(t, a1.LefterThan(Degrees(-181).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(-180).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(-179).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-90).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(-1).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(0).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(1).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(10).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(45).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(90).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(135).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(179).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(180).ToRadians()))
		assert.False(t, a1.LefterThan(Degrees(181).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(235).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(275).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(300).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(330).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(359).ToRadians()))
		assert.True(t, a1.LefterThan(Degrees(360).ToRadians()))
	})
}
