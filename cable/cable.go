// Copyright (c) 2017-2022 The emc developers. All rights reserved.
// Project site: https://github.com/gotmc/emc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package cable

import "gonum.org/v1/gonum/floats"

type CableType string

const (
	RG58   CableType = "RG-58"
	RG58U  CableType = "RG-58/U"
	LMR195 CableType = "LMR-195"
	LMR400 CableType = "LMR-400"
)

func (c CableType) String() string {
	return string(c)
}

// Loss returns the frequencies and cable loss based on the cable type and
// length in feet.
func Loss(cableType CableType, length float64) ([]float64, []float64) {
	defaultLength := 0.0
	freqs := []float64{}
	loss := []float64{}

	switch cableType {
	case RG58:
		defaultLength = 100.0
		freqs = []float64{1.0e6, 1.0e7, 5.0e7, 1.0e8, 2.0e8, 4.0e8, 7.0e8, 9.0e8,
			1.0e9}
		loss = []float64{0.4, 1.4, 3.3, 4.9, 7.3, 11.2, 16.9, 20.1, 21.5}
	case RG58U:
		defaultLength = 100.0
		freqs = []float64{9.0e3, 1.0e4, 2.5e4, 5.0e4, 1.0e5, 5.0e5, 1.0e6, 5.0e6,
			1.0e7, 5.0e7, 7.5e7, 1.0e8, 1.5e8, 2.0e8, 3.0e8, 4.0e8, 5.0e8, 6.0e8,
			7.0e8, 8.0e8, 9.0e8, 1.0e9}
		loss = []float64{0.0384, 0.0404, 0.0640, 0.0906, 0.1283, 0.2873, 0.4067,
			0.9111, 1.2896, 3.0000, 3.5403, 4.4, 5.0109, 6.00, 7.0924, 8.5, 9.1619,
			10.0385, 10.8448, 11.5955, 13.0, 13.0}
	case LMR195:
		defaultLength = 100.0
		freqs = []float64{3.0e5, 3.0e6, 3.0e7, 5.0e7, 1.5e8, 2.2e8, 4.5e8, 9.0e8,
			1.5e9}
		loss = []float64{0.20, 0.60, 2.00, 2.50, 4.40, 5.40, 7.80, 11.10, 14.50}
	case LMR400:
		defaultLength = 100.0
		freqs = []float64{9.0e3, 1.0e4, 2.5e4, 5.0e4, 1.0e5, 5.0e5, 1.0e6, 5.0e6,
			1.0e7, 3.0e7, 5.0e7, 7.5e7, 1.0e8, 1.5e8, 2.0e8, 2.2e8, 3.0e8, 4.0e8,
			4.5e8, 5.0e8, 6.0e8, 7.0e8, 8.0e8, 9.0e8, 1.0e9, 1.5e9}
		loss = []float64{0.0112, 0.0118, 0.0188, 0.0267, 0.0380, 0.0860,
			0.1223, 0.2771, 0.3490, 0.7000, 0.900, 1.0968, 1.2695, 1.50, 1.8054, 1.90,
			2.2185, 2.5767, 2.70, 2.8759, 3.155, 3.4121, 3.6516, 3.9, 4.09, 5.1}
	}
	// Adjust the loss based on given length.
	floats.Scale(length/defaultLength, loss)
	return freqs, loss
}
