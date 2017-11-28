package lang

import (
	"io/ioutil"
	"log"
)

const (
	//定义提醒的层级关系
	debu     = "inf"
	inf      = "inf"
	notic    = "inf"
	warnin   = "warnin"
	erro     = "dange"
	critica  = "dange"
	aler     = "dange"
	emergenc = "dange"
	processe = "inf"

	//定义图像显示层级
	debug     = "info"
	info      = "info"
	notice    = "info"
	warning   = "warning"
	error     = "warning"
	critical  = "warning"
	alert     = "warning"
	emergency = "warning"
	processed = "inf"

	MaxFileSize = 1 << 20
)

func All() {
	var log []string
	pattern := `\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}([\+-]\d{4})?\].*`
	fileColl := GetFiles("./")
}

//GetFiles return all files by give the basename
func GetFiles(path string) []string {
	fileList, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}
	var fileCollec []string
	for _, file := range fileList {
		if !file.IsDir() {
			fileCollec = append(fileCollec, file.Name())
		}
	}
	return fileCollec
}
