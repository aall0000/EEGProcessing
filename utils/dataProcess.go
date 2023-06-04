package utils

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"sort"
)

func NumberPlotERD(data []float64, lineName string, savePath string) error {
	points := plotter.XYs{}
	for i := 0; i < len(data); i++ {
		points = append(points, plotter.XY{
			X: float64(i),
			Y: data[i],
		})
	}
	plt := plot.New()
	arr := make([]float64, len(data))
	copy(arr, data)
	sort.Float64s(arr)
	yMax := arr[len(arr)-1]
	yMin := arr[0]
	xMax := float64(len(data))
	plt.Y.Max, plt.Y.Min, plt.X.Max, plt.X.Min = yMax, yMin, xMax, 0
	err := plotutil.AddLines(plt, lineName, points)
	if err != nil {
		return err
	}
	err = plt.Save(5*vg.Inch, 5*vg.Inch, savePath)
	if err != nil {
		return err
	}
	return nil
}

func NumberPlotPSD(data []float64, bands []string, epochs int, savePath string) error {
	var err error
	plt := plot.New()
	arr := make([]float64, len(data))
	copy(arr, data)
	sort.Float64s(arr)
	yMax := arr[len(arr)-1]
	yMin := arr[0]
	xMax := float64(epochs)
	plt.Y.Max, plt.Y.Min, plt.X.Max, plt.X.Min = yMax, yMin, xMax, 0
	lines := make([]plotter.XYs, len(bands))
	for i := 0; i < len(bands); i++ {
		rhythmData := data[i*epochs : (i+1)*epochs]
		//points := plotter.XYs{}
		for j := 0; j < epochs; j++ {
			lines[i] = append(lines[i], plotter.XY{
				X: float64(j),
				Y: rhythmData[j],
			})
		}
	}
	switch len(bands) {
	case 1:
		err = plotutil.AddLines(plt, bands[0], lines[0])
		if err != nil {
			return err
		}
	case 2:
		err := plotutil.AddLines(plt,
			bands[0], lines[0],
			bands[1], lines[1])
		if err != nil {
			return err
		}
	case 3:
		err := plotutil.AddLines(plt,
			bands[0], lines[0],
			bands[1], lines[1],
			bands[2], lines[2],
		)
		if err != nil {
			return err
		}
	case 4:
		err := plotutil.AddLines(plt,
			bands[0], lines[0],
			bands[1], lines[1],
			bands[2], lines[2],
			bands[3], lines[3],
		)
		if err != nil {
			return err
		}
	}

	err = plt.Save(5*vg.Inch, 5*vg.Inch, savePath)
	if err != nil {
		return err
	}
	return nil
}

func ConvertToExcelERD(data []float64, filedName1, filedName2, savePath string, digitNum int32) (*excelize.File, error) {
	var err error
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	err = f.SetCellValue("Sheet1", "A1", filedName1)
	if err != nil {
		return nil, err
	}
	err = f.SetCellValue("Sheet1", "B1", filedName2)
	if err != nil {
		return nil, err
	}
	res := make([]float64, len(data))
	for i, v := range data {
		value, _ := decimal.NewFromFloat(v).Round(digitNum).Float64()
		res[i] = value
	}

	for i, v := range res {
		err = f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), i)
		if err != nil {
			return nil, err
		}
		err = f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), v)
		if err != nil {
			return nil, err
		}
	}

	err = f.SaveAs(savePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func ConvertToExcelPSD(data []float64, bands []string, epochs int, savePath string, digitNum int32) (*excelize.File, error) {
	var err error
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	cellVal := 'A'
	length := len(bands)
	res := make([]float64, len(data))
	for i, v := range data {
		value, _ := decimal.NewFromFloat(v).Round(digitNum).Float64()
		res[i] = value
	}
	for i := 0; i < length; i++ {
		cellName := string(cellVal)
		cellVal++
		err = f.SetCellValue("Sheet1", fmt.Sprintf("%s1", cellName), bands[i])
		if err != nil {
			return nil, err
		}
		rhythmData := res[i*epochs : (i+1)*epochs]
		for j := 0; j < epochs; j++ {
			err = f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", cellName, j+2), rhythmData[j])
			if err != nil {
				return nil, err
			}
		}
	}
	err = f.SaveAs(savePath)
	if err != nil {
		return nil, err
	}
	return f, nil
}
