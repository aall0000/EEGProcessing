package utils

import "github.com/pkg/errors"

func Smooth(data []float64, smoothPara int) ([]float64, error) {
	if smoothPara < 0 {
		return nil, errors.Errorf("smoothPara must be positive")
	}
	sData := make([]float64, len(data))
	b := (smoothPara - 1) / 2
	if len(data) < smoothPara {
		for i := 0; i < len(data); i++ {
			if i < len(data)/2 {
				for j := -i; j < i+1; j++ {
					sData[i] += data[i+j]
				}
				sData[i] = sData[i] / float64(i*2+1)
			} else {
				k := len(data) - 1 - i
				for j := -k; j < k+1; j++ {
					sData[i] += data[i+j]
				}
				sData[i] = sData[i] / float64(k*2+1)
			}
		}
	} else {
		for i := 0; i < len(data); i++ {
			if i < b {
				for j := -i; j < i+1; j++ {
					sData[i] += data[i+j]
				}
				sData[i] = sData[i] / float64(2*i+1)
			} else if (i > b || i == b) && len(data)-i > b {
				for j := -b; j < b+1; j++ {
					sData[i] += data[i+j]
				}
				if smoothPara%2 == 0 {
					sData[i] = sData[i] / float64(smoothPara-1)
				} else {
					sData[i] = sData[i] / float64(smoothPara)
				}
			} else if i > b && len(data)-i <= b {
				k := len(data) - 1 - i
				for j := -k; j < k+1; j++ {
					sData[i] += data[i+j]
				}
				sData[i] = sData[i] / float64(2*k+1)
			}
		}
	}
	return sData, nil
}

//func DurationSum(start int, end int, arr []float64) float64 {
//	var sum float64 = 0
//	for i := start; i <= end; i++ {
//		sum += arr[i]
//	}
//	return sum
//}
