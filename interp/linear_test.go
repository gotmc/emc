package interp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinearInterpolation(t *testing.T) {
	testCases := []struct {
		y1       []float64
		x1       []float64
		x2       []float64
		expected []float64
	}{
		{
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{1.5, 2.5, 3.5},
			[]float64{1.5, 2.5, 3.5},
		},
	}
	// Loop through each testCases
	for _, tc := range testCases {
		y2 := Linear(tc.y1, tc.x1, tc.x2)
		assert.Equal(t, y2, tc.expected)
	}
}
