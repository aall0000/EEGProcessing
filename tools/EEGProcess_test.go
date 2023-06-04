package tools

import (
	"EEG/models"
	"testing"
)

func BenchmarkEEGProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sys := models.SysParams{
			SmoothParam: 150,
			RStart:      0,
			REnd:        1,
		}
		element := ElementRead("./files/test.mat")
		eegData := ElementParse(element, 1000)
		erdOp := CreateERDOp(sys, *eegData)
		_, _ = erdOp.Operate()
	}
}
