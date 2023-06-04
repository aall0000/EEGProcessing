package tools

import "EEG/models"

type EEGData struct {
	Data       []float64 //原数据集
	SampleRate int       //采样率
	Epochs     int       //事件数
	Frames     int       //每个事件时长
}

func CreateEEGData(frames, epochs, sampleRate int, data []float64) *EEGData {
	return &EEGData{
		Frames:     frames,
		Epochs:     epochs,
		Data:       data,
		SampleRate: sampleRate,
	}
}

type EEGProcess interface {
	SetData([]float64)
	SetSampleRate(int)
	Operate() ([]float64, error)
}

func (ed *EEGData) SetData(data []float64) {
	ed.Data = data
}

func (ed *EEGData) SetSampleRate(sr int) {
	ed.SampleRate = sr
}

func (ed *EEGData) SetEpochs(ep int) {
	ed.Epochs = ep
}

func (ed *EEGData) SetFrames(fr int) {
	ed.Frames = fr
}

type ERDOperation struct {
	SmoothPara int
	RStart     float32 //参考时间段开始时间
	REnd       float32 //参考时间段结束时间
	EEGData
}

func CreateERDOp(sysParam models.SysParams, eegData EEGData) ERDOperation {
	return ERDOperation{
		SmoothPara: sysParam.SmoothParam,
		RStart:     sysParam.RStart,
		REnd:       sysParam.REnd,
		EEGData:    eegData,
	}
}

func (eo *ERDOperation) SetSmoothPara(sp int) {
	eo.SmoothPara = sp
}

func (eo *ERDOperation) SetRStart(rs float32) {
	eo.RStart = rs
}

func (eo *ERDOperation) SetREnd(re float32) {
	eo.REnd = re
}

func (eo *ERDOperation) Operate() ([]float64, error) {
	erdOp := &ERDOperation{
		SmoothPara: eo.SmoothPara,
		RStart:     eo.RStart,
		REnd:       eo.REnd,
		EEGData:    eo.EEGData,
	}
	res, err := ERDCompute(erdOp)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type PSDOperation struct {
	RhythmOptions []string
	EEGData
}

func CreatePSDOp(sysParam models.SysParams, eegData EEGData) PSDOperation {
	return PSDOperation{
		RhythmOptions: sysParam.BandOptions,
		EEGData:       eegData,
	}
}

func (po *PSDOperation) SetBandOption(bo []string) {
	po.RhythmOptions = bo
}

func (po *PSDOperation) Operate() ([]float64, error) {
	psdOp := &PSDOperation{
		RhythmOptions: po.RhythmOptions,
		EEGData:       po.EEGData,
	}
	res, err := PSDCompute(psdOp)
	if err != nil {
		return nil, err
	}
	return res, nil
}
