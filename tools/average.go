package tools

// GetAverage calculates the average of a list of integers and returns the rounded result
func GetAverage(numList []float64) float64 {
	sum := 0.0
	for _, n := range numList {
		sum += float64(n)
	}
	result := float64(sum) / float64(len(numList))
	return result
}
