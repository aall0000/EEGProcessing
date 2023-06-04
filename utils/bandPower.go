package utils

import "github.com/pkg/errors"

func BandPower(pxx []float64, f []float64, fStart, fEnd float64) (float64, error) {
	if fStart >= fEnd {
		return 0, errors.Errorf("fStart must smaller than fEnd")
	}
	if len(pxx) != len(f) {
		return 0, errors.Errorf("len(pxx) must be equal to len(f)")
	}
	if !DiffPositive(f) {
		return 0, errors.Errorf("f is not an increasing sequence")
	}
	idx1 := LastOneLessOREqual(len(f), func(i int) bool {
		return f[i] <= fStart
	})
	idx2 := FirstOneGreaterOREqual(len(f), func(i int) bool {
		return f[i] >= fEnd
	})
	width := DiffArray(f)
	power := 0.0
	for i := idx1; i < idx2+1; i++ {
		power += width[i] * pxx[i]
	}
	return power, nil
}
