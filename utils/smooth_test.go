package utils

import (
	"testing"
)

func TestSmooth(t *testing.T) {
	cases := []struct {
		Name        string
		Arr         []float64
		SmoothPara  int
		ArrExpected []float64
	}{
		{"smoothParaTooBig", []float64{2, 2, 1, 3, 4, 5}, 9, []float64{2, 5.0 / 3, 2.4, 3, 4, 5}},
		{"smoothParaNormal", []float64{2, 2, 1, 3, 4, 5}, 5, []float64{2, 5.0 / 3, 2.4, 3, 4, 5}},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans, _ := Smooth(c.Arr, c.SmoothPara); !sliceEqual(ans, c.ArrExpected) {
				t.Errorf("expected: %v ,got: %v",
					c.ArrExpected, ans)
			}
		})
	}
}

func sliceEqual(a []float64, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
