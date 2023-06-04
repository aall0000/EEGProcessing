package main

import (
	"EEG/models"
	"EEG/routers"
	"EEG/tools"
)

func main() {
	//定期清理files文件夹
	go tools.RegularEmpty()
	//实现web功能
	sysParam := models.CreateNew()
	var res []float64
	r := routers.SetupRouter(sysParam, res)
	r.Run(":8081")

}
