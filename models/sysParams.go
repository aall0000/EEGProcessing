package models

type SysParams struct {
	FileDst     string
	SampleRate  int
	SmoothParam int
	RStart      float32
	REnd        float32
	BandOptions []string
}

//var Result []float64

func CreateNew() *SysParams {
	return &SysParams{}
}
