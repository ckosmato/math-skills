package tools

import "math-skills/subtools"

// GetStandardDeviation calculates the standard deviation of a list of integers
func GetStandardDeviation(numList []float64) int {
	vrncTemp := GetVariance(numList)

	vrnc := subtools.FindSquare(vrncTemp)

	return vrnc
}
