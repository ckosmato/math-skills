package tools

import (
	"math-skills/subtools"
)

// GetMedian calculates the median of a sorted list of integers
func GetMedian(numList []float64) int {
	sorted := subtools.GetSorted(numList)

	if len(sorted)%2 == 0 {
		return subtools.Rounded((float64(sorted[(len(sorted)-1)/2]) + float64((sorted[len(sorted)/2]))) / float64(2))
	}

	return int(sorted[len(sorted)/2])
}
