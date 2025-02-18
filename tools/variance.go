package tools

// GetVariance calculates the variance of a list of integers
func GetVariance(numList []float64) float64 {
	avrg := GetAverage(numList)
	vrncList := []float64{}
	final := []float64{}

	for _, n := range numList {
		vrncList = append(vrncList, (avrg-float64(n))*(avrg-float64(n)))
	}
	for _, f := range vrncList {
		final = append(final, f)
	}
	return GetAverage(final)
}
