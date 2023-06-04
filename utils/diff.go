package utils

func DiffPositive(w []float64) bool {
	for i := 0; i < len(w)-1; i++ {
		if w[i+1]-w[i] < 0 {
			return false
		}
	}
	return true
}

func DiffArray(w []float64) []float64 {
	width := make([]float64, len(w))
	for i := 0; i < len(w)-1; i++ {
		width[i] = w[i+1] - w[i]
	}
	return width
}
