package subtools

// FindSquare takes a  n number and returns the largest number whose square is less than or equal to that n number.

func FindSquare(n float64) int {
	sq := 0.0
	for i := 0.0; i*i <= n; i += 0.01 {
		sq = i
	}
	return Rounded(sq)
}
