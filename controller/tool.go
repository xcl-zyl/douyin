/*
工具包：需要用到的功能函数
tools：provide some useful func
*/

package controller

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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
	return fileNames[:]
}

// 查询用户是否存在以及查询密码是否正确
// judge the user is exist and judge the password is correct
func GetIsExist(username string, password ...string) (int, string) {
	userId := 0
	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
	if err != nil {
		fmt.Println("数据库连接失败" + err.Error())
		return userId, username
	}

	if len(password) != 0 { // select 查了多少元素 scan 就必须读多少元素， 个数必须统一
		Sql := fmt.Sprintf("select userId, userName, password from user where userName = '%s' and password = '%s'", username, password[0])
		db.QueryRow(Sql).Scan(&userId, &username, &password)
	} else {
		Sql := fmt.Sprintf("select userId, userName from user where userName = '%s'", username)
		db.QueryRow(Sql).Scan(&userId, &username)
	}
	defer db.Close()
	return userId, username
}

// 向数据库插入用户信息
// insert user in mysql
func AddUser(username, password string) {
	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("insert into user values (0, '%s', '%s')", username, password)
	db.Exec(Sql)
	defer db.Close()
}
