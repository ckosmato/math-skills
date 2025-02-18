package subtools

// Rounded rounds a number to the nearest integer number

func Rounded(rounded float64) int {
	rounded = rounded * 10
	if int(rounded)%10 < 5 {
		return int(rounded) / 10
	}
	return (int(rounded) / 10) + 1
}
