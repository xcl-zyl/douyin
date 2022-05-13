/*
这个文件用于测试功能函数，在使用前进行简单测试，测试后的函数放入核心代码中调用
this file is to test func, using the func after test it.
*/

package test

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net"
	"strings"

	_ "github.com/go-sql-driver/mysql"
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

func GetIsExist(username string, password ...string) (int, string) {
	userId := 0
	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
	if err != nil {
		fmt.Println("数据库连接失败")
		return userId, username
	}

	if len(password) != 0 {
		// select 查了多少元素 scan 就必须读多少元素， 个数必须统一
		Sql := fmt.Sprintf("select userId, userName, password from user where userName = '%s' and password = '%s'", username, password[0])
		db.QueryRow(Sql).Scan(&userId, &username, &password)
	} else {
		Sql := fmt.Sprintf("select userId, userName from user where userName = '%s'", username)
		db.QueryRow(Sql).Scan(&userId, &username)
	}
	// fmt.Println(username, password)
	// if Err != sql.ErrNoRows {
	// 	res = userid
	// }
	defer db.Close()
	return userId, username
}

func GetUser(username string) (int, string) {
	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
	if err != nil {
		fmt.Println("数据库连接失败")
		return 0, ""
	}
	var userId int
	var userName string
	Sql := fmt.Sprintf("select userId, userName from user where userName = '%s'", username)
	db.QueryRow(Sql).Scan(&userId, &userName)
	defer db.Close()
	return userId, userName
}

func AddUser(username, password string) {
	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("insert into user values (0, '%s', '%s')", username, password)
	db.Exec(Sql)
	// result, _ := db.Exec(sql)
	// n, _ := result.RowsAffected()
	// fmt.Println("受影响的记录数是", n)
	defer db.Close()
}

func AddVideo(author, playUrl, coverUrl string) {
	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("insert into video values (0, '%s', '%s', '%s')", author, playUrl, coverUrl)
	db.Exec(Sql)
	defer db.Close()
}

// func GetVideo() []string {
// 	var res []Video
// 	db, err := sql.Open("mysql", "xcl:xcl201314@(localhost:3306)/douyin")
// 	if err != nil {
// 		fmt.Println("数据库连接失败")
// 		return res
// 	}
// 	rows, _ := db.Query("select * from video")
// 	var id int64
// 	var author, playUrl, coverUrl string
// 	for rows.Next() {
// 		rows.Scan(&id, &author, &playUrl, &coverUrl)
// 		fmt.Println(id, "--", author, playUrl, coverUrl)
// 	}
// 	defer db.Close()
// }

func TestSlice() {
	s := []int{1, 2, 3}
	a := s[0:30]
	println(a)
}
