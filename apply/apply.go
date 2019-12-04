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
