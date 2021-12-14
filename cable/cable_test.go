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
		cable  CableType
		length float64
		want   []float64
	}{
		{
			cable:  LMR400,
			length: 50.0,
			want:   []float64{0.0056, 0.0059},
		},
	}
	for _, test := range tests {
		name := fmt.Sprintf("%s at %f.2 ft", test.cable, test.length)
		t.Run(name, func(t *testing.T) {
			_, gotLoss := Loss(test.cable, test.length)
			assertFloat64(t, name, gotLoss[0], test.want[0], 0.1e-6)
		})
	}
}

func assert(t *testing.T, label string, got, want interface{}) {
	if got != want {
		t.Errorf("\ngot  = `%#v` for %s\nwant = `%#v`", got, label, want)
	}
}

func assertFloat64(t *testing.T, label string, got, want, tolerance float64) {
	if diff := math.Abs(want - got); diff >= tolerance {
		t.Errorf("\ngot %s  = %#v \t\nwant %s = %#v", label, got, label, want)
	}
}
