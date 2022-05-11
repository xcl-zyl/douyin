package controller

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

// 用于获取主机外网ip
// this func to get host net ip
func GetHostIp() string {
	// 通过udp方式访问其它网址获得外网ip
	// use udp connect other net to get host ip
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	// 返回前去掉端口
	// remove port before return
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

// 用于获取指定路径下所有文件名，不包括子文件夹
// to get filename in the path, but not include subfolder
func GetAllFile(path string) []string {
	var fileNames []string
	rd, err := ioutil.ReadDir(path)
	if err != nil {
		// 目标文件夹不存在时，返回空字符串组
		// destination folder is not exist, return null []string
		fmt.Println("read dir fail:", err)
		return fileNames
	}

	for _, fi := range rd { // _代表序号，fi表示值。 _ represent array number, fi represent array value.
		if !fi.IsDir() {
			fileName := fi.Name()
			fileNames = append(fileNames, fileName)
		}
	}
	return fileNames
}
