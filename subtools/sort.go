package subtools

// GetSorted performs a simple bubble sort on a slice of integers.

func GetSorted(numList []float64) []float64 {
	for i := 0; i < len(numList)-1; i++ {
		if numList[i] > numList[i+1] {

			tempNum := numList[i]
			numList = append(numList[:i], numList[i+1:]...)
			numList = append(numList, tempNum)
			i = -1
		}
	}

	return numList
}
