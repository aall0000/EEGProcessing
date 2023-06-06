package routers

import (
	"EEG/models"
	"EEG/tools"
	"EEG/utils"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func SetupRouter(sys *models.SysParams, res []float64) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	//获取文件
	router.MaxMultipartMemory = 8 << 20 //8Mb
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		dst := "./files/" + file.Filename
		sys.FileDst = dst
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//路由嵌套
	erdGroup := router.Group("/erd")
	{
		erdGroup.POST("/compute", func(c *gin.Context) {
			sys.SampleRate, _ = strconv.Atoi(c.PostForm("sample_rate"))
			sys.SmoothParam, _ = strconv.Atoi(c.PostForm("smooth_param"))
			val1, _ := strconv.ParseFloat(c.PostForm("r_start"), 32)
			sys.RStart = float32(val1)
			val2, _ := strconv.ParseFloat(c.PostForm("r_end"), 32)
			sys.REnd = float32(val2)
			//运行ERD
			sys.FileDst = "D:\\360MoveData\\Users\\Hu\\Desktop\\gotest\\特征提取工具\\EEG\\files\\test.mat"
			element := tools.ElementRead(sys.FileDst)
			eegData := tools.ElementParse(element, sys.SampleRate)
			erdOp := tools.CreateERDOp(*sys, *eegData)
			res, _ = erdOp.Operate()
			//fmt.Println(res[:5])
			savePath := "D:\\360MoveData\\Users\\Hu\\Desktop\\gotest\\奥集能前端\\public\\matfile\\ERDTest.png"
			utils.NumberPlotERD(res, "line1", savePath)
			//流形式传给前端
			c.JSON(http.StatusOK, gin.H{
				"epochs": eegData.Epochs,
				"frames": eegData.Frames,
			})
		})
		erdGroup.GET("/view", func(c *gin.Context) {
			savePath := "./files/test.png"
			utils.NumberPlotERD(res, "line1", savePath)
			c.Header("response-type", "blob")
			srcByte, _ := ioutil.ReadFile(savePath)
			//流形式传给前端
			c.Data(http.StatusOK, "image/png", srcByte)
		})
		erdGroup.GET("/xlsx", func(c *gin.Context) {
			//fmt.Println(res)
			//生成ERD的excel表
			savePath := "./files/test1.xlsx"
			file, err := utils.ConvertToExcelERD(res, "Time", "ERDValue", savePath, 3)
			if err != nil {
				log.Fatal(err)
			}
			//将生成的excel发给前端
			c.Header("response-type", "blob")
			data, _ := file.WriteToBuffer()
			c.Data(http.StatusOK, "application/vnd.ms-excel", data.Bytes())
		})
	}

	//路由嵌套
	psdGroup := router.Group("/psd")
	{
		var epochs int
		psdGroup.POST("/compute", func(c *gin.Context) {
			sys.BandOptions = c.PostFormArray("band_options[]")
			sys.SampleRate, _ = strconv.Atoi(c.PostForm("sample_rate"))
			sys.FileDst = "D:\\360MoveData\\Users\\Hu\\Desktop\\gotest\\特征提取工具\\EEG\\files\\test.mat"
			element := tools.ElementRead(sys.FileDst)
			eegData := tools.ElementParse(element, sys.SampleRate)
			epochs = eegData.Epochs
			psdOp := tools.CreatePSDOp(*sys, *eegData)
			res, _ = psdOp.Operate()
			savePath := "D:\\360MoveData\\Users\\Hu\\Desktop\\gotest\\奥集能前端\\public\\matfile\\PSDTest.png"
			utils.NumberPlotPSD(res, sys.BandOptions, epochs, savePath)
			c.JSON(http.StatusOK, gin.H{
				"epochs": epochs,
				"frames": eegData.Frames,
			})
		})
		psdGroup.GET("/view", func(c *gin.Context) {
			savePath := "./files/PSDTest.png"
			utils.NumberPlotPSD(res, sys.BandOptions, epochs, savePath)
			c.Header("response-type", "blob")
			srcByte, _ := ioutil.ReadFile(savePath)
			//流形式传给前端
			c.Data(http.StatusOK, "image/png", srcByte)
		})
		psdGroup.GET("/xlsx", func(c *gin.Context) {
			savePath := "./files/PSDTest.xlsx"
			file, err := utils.ConvertToExcelPSD(res, sys.BandOptions, epochs, savePath, 3)
			if err != nil {
				log.Fatal(err)
			}
			c.Header("response-type", "blob")
			data, _ := file.WriteToBuffer()
			c.Data(http.StatusOK, "application/vnd.ms-excel", data.Bytes())
		})
	}

	return router
}
