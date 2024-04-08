// Copyright (c) 2017-2024 The emc developers. All rights reserved.
// Project site: https://github.com/gotmc/emc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

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
			[]float64{0.5, 1.5, 2.5, 3.5, 4.5},
			[]float64{0.5, 1.5, 2.5, 3.5, 4.5},
		},
		{
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{1.0, 1.1, 1.2, 1.3, 1.4},
			[]float64{1.0, 1.1, 1.2, 1.3, 1.4},
		},
	}
	// Loop through the testCases
	for _, tc := range testCases {
		y2 := Linear(tc.y1, tc.x1, tc.x2)
		assert.Equal(t, y2, tc.expected)
	}
}

func TestLocatingSubintervalIndex(t *testing.T) {
	testCases := []struct {
		x        []float64
		z        float64
		expected int
	}{
		{[]float64{1.0, 2.0, 3.0, 4.0}, 1.5, 0},
		{[]float64{1.0, 2.0, 3.0, 4.0}, 1.0, 0},
		{[]float64{1.0, 2.0, 3.0, 4.0}, 0.5, 0},
		{[]float64{1.0, 2.0, 3.0, 4.0}, 3.5, 2},
		{[]float64{1.0, 2.0, 3.0, 4.0}, 4.5, 2},
	}
	// Loop through the testCases
	for _, tc := range testCases {
		i := subintervalIndex(tc.x, tc.z)
		assert.Equal(t, i, tc.expected)
	}
}

func TestLocatingSubintervalIndices(t *testing.T) {
	testCases := []struct {
		x        []float64
		z        []float64
		expected []int
	}{
		{[]float64{1.0, 2.0, 3.0, 4.0}, []float64{1.5, 2.5}, []int{0, 1}},
		{[]float64{1.0, 2.0, 3.0, 4.0}, []float64{1.5, 3.5}, []int{0, 2}},
	}
	// Loop through each testCases
	for _, tc := range testCases {
		i := subintervalIndices(tc.x, tc.z)
		assert.Equal(t, i, tc.expected)
	}
}

func TestCalculatingSlopes(t *testing.T) {
	testCases := []struct {
		x        []float64
		y        []float64
		expected []float64
	}{
		{
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{1.0, 1.0, 1.0},
		},
		{
			[]float64{1.0, 2.0, 3.0, 4.0},
			[]float64{2.0, 4.0, 6.0, 7.0},
			[]float64{2.0, 2.0, 1.0},
		},
	}
	// Loop through each testCases
	for _, tc := range testCases {
		calculated := slopes(tc.x, tc.y)
		assert.Equal(t, calculated, tc.expected)
	}
}
