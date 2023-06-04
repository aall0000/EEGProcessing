package tools

import (
	"github.com/florianl/matf"
	"io"
	"log"
	"reflect"
)

func ElementRead(matPath string) matf.MatMatrix {
	//读取matlab的头文件
	modelFile, err := matf.Open(matPath)
	if err != nil {
		log.Fatal(err)
	}
	defer matf.Close(modelFile)

	//读取数据
	element, err := matf.ReadDataElement(modelFile)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	return element
}

func ElementParse(elem matf.MatMatrix, sampleRate int) *EEGData {
	channels, frames, epochs, err := elem.Dimensions()
	if channels != 1 {
		log.Fatal("the number of channel should be 1")
	}
	if err != nil {
		log.Fatal(err)
	}
	var data []float64
	slice := reflect.ValueOf(elem.Content.(matf.NumPrt).RealPart)
	for i := 0; i < slice.Len(); i++ {
		value := reflect.ValueOf(slice.Index(i).Interface()).Float()
		data = append(data, value)
	}
	eegData := CreateEEGData(frames, epochs, sampleRate, data)
	return eegData
}
