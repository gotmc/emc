// Copyright (c) 2017-2022 The emc developers. All rights reserved.
// Project site: https://github.com/gotmc/emc
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package apply

import "github.com/gotmc/emc/interp"

// AntennaFactor linearly interpolates the antennaFactors given the
// antennaFreqs to match the analyzerFreqs. The given antennaFactors and
// analyzerReadings should be in dB and dBuV, respectively. The returned slice
// contains the incident electric field in dBuV/m.
func AntennaFactor(analyzerReadings, analyzerFreqs, antennaFactors, antennaFreqs []float64) []float64 {
	antennaFactorsAtAnalyzerFreqs := interp.Linear(antennaFactors, antennaFreqs, analyzerFreqs)
	return antennaFactorsAtAnalyzerFreqs
}
