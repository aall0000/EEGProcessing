package tools

import (
	"EEG/utils"
	"github.com/mjibson/go-dsp/spectral"
	"github.com/mjibson/go-dsp/window"
	"github.com/pkg/errors"
)

func PSDCompute(psdOp *PSDOperation) ([]float64, error) {
	pos := &spectral.PwelchOptions{
		NFFT: 512,
		Window: func(i int) []float64 {
			return window.Hamming(512)
		},
		Noverlap: 256,
	}
	var res []float64
	fStart, fEnd := 0.0, 0.0
	for _, option := range psdOp.RhythmOptions {
		switch option {
		case "delta":
			fStart, fEnd = 0.5, 4
		case "theta":
			fStart, fEnd = 4, 8
		case "alpha":
			fStart, fEnd = 8, 13
		case "beta":
			fStart, fEnd = 13, 30
		default:
			err := errors.Errorf("rhythm option got wrong value")
			if err != nil {
				return nil, err
			}
		}

		for i := 0; i < psdOp.Epochs; i++ {
			tData := psdOp.Data[i*psdOp.Frames : (i+1)*psdOp.Frames]
			power, f := spectral.Pwelch(tData, float64(psdOp.SampleRate), pos)
			pxx := power
			rhythm, err := utils.BandPower(pxx, f, fStart, fEnd)
			if err != nil {
				return nil, err
			}
			res = append(res, rhythm)
		}
	}
	return res, nil
}
