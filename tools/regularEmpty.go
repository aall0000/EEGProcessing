package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"
)

func RegularEmpty() {
	t := time.NewTicker(time.Hour * 168)
	go executeJob(t)
	select {} //阻塞函数退出
}

func executeJob(t interface{}) {
	sourceUrl := "./files"
	println(sourceUrl)
	for {
		select {
		case <-t.(*time.Ticker).C:
			fmt.Println("执行任务-清空files文件夹")
			deleteFileByDir(sourceUrl)
		}
	}
}

func deleteFileByDir(url string) {
	dir, _ := ioutil.ReadDir(url)
	for _, d := range dir {
		fmt.Println("删除：", d.Name())
		os.RemoveAll(path.Join([]string{"files", d.Name()}...))
	}
}
