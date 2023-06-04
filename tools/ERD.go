package tools

import (
	"EEG/utils"
)

func ERDCompute(erdOp *ERDOperation) ([]float64, error) {
	doubleData := make([]float64, len(erdOp.Data))
	dProc := make([]float64, erdOp.Frames)
	for i := range erdOp.Data {
		doubleData[i] = erdOp.Data[i] * erdOp.Data[i]
	}
	for i := 0; i < erdOp.Frames; i++ {
		for j := 0; j < erdOp.Epochs; j++ {
			dProc[i] += doubleData[j*erdOp.Frames+i]
		}
		dProc[i] = dProc[i] / float64(erdOp.Epochs)
	}
	res, err := utils.Smooth(dProc, erdOp.SmoothPara)
	if err != nil {
		return nil, err
	}

	RStartTime := int(erdOp.RStart * float32(erdOp.SampleRate))
	REndTime := int(erdOp.REnd * float32(erdOp.SampleRate))
	RTime := REndTime - RStartTime
	Rdata := make([]float64, RTime)
	copy(Rdata, res[RStartTime:REndTime])
	var R float64
	for i := 0; i < RTime; i++ {
		R += Rdata[i]
	}
	R = R / float64(RTime)

	for i := range res {
		res[i] = (R - res[i]) / R
	}
	return res, nil
}
