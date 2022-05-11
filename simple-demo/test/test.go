// 这个文件用于测试功能函数，在使用前进行简单测试，测试后的函数放入核心代码中调用
// this file is to test func, using the func after test it.

package test

import (
	"fmt"
	"io/ioutil"
	"net"
	"strings"
)

func Test() string {
	conn, err := net.Dial("udp", "baidu.com:80")
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer conn.Close()
	fmt.Println(conn.LocalAddr().String())
	return strings.Split(conn.LocalAddr().String(), ":")[0]
}

func GetAllFile(pathname string) []string {
	var s []string
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		fmt.Println("read dir fail:", err)
		return s
	}

	for _, fi := range rd {
		if !fi.IsDir() {
			fullName := fi.Name()
			s = append(s, fullName)
		}
	}
	return s
}
