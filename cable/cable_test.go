// Copyright (c) 2017-2022 The emc developers. All rights reserved.
// Project site: https://github.com/gotmc/emc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package cable

import (
	"fmt"
	"math"
	"testing"
)

func TestLoss(t *testing.T) {
	var tests = []struct {
		cable    CableType
		length   float64
		wantFreq []float64
		wantLoss []float64
	}{
		{
			cable:    LMR400,
			length:   50.0,
			wantLoss: []float64{0.0056, 0.0059},
		},
		{
			cable:  RG58,
			length: 100.0,
			wantFreq: []float64{
				1.00e6, 1.00e7, 5.00e7, 1.00e8, 2.00e8, 4.00e8, 7.00e8, 9.00e8, 1.00e9,
			},
			wantLoss: []float64{0.4, 1.4, 3.3, 4.9, 7.3, 11.2, 16.9, 20.1, 21.5},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%s at %f.2 ft", test.cable, test.length)
		t.Run(name, func(t *testing.T) {
			_, gotLoss := Loss(test.cable, test.length)
			assertFloat64(t, name, gotLoss[0], test.wantLoss[0], 0.1e-6)
		})
	}
}

func assert(tb testing.TB, label string, got, want interface{}) {
	tb.Helper()
	if got != want {
		tb.Errorf("\ngot  = `%#v` for %s\nwant = `%#v`", got, label, want)
	}
}

func assertFloat64(tb testing.TB, label string, got, want, tolerance float64) {
	tb.Helper()
	if diff := math.Abs(want - got); diff >= tolerance {
		tb.Errorf("\ngot %s  = %#v \t\nwant %s = %#v", label, got, label, want)
	}
}
