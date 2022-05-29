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

const dbName string = "mysql"
const dbConnect string = "xcl:xcl201314@(localhost:3306)/douyin"

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
func GetIsExist(username string, password ...string) (int64, string) {
	userId := int64(0)
	db, err := sql.Open(dbName, dbConnect)
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
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("insert into user values (0, '%s', '%s')", username, password)
	db.Exec(Sql)
	defer db.Close()
}

// 插入视频信息，插入视频时默认未点赞，评论数点赞数为0
// insert video inf
func AddVideo(author, playUrl, coverUrl string) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("insert into video values (0, '%s', '%s', '%s', '%d', '%d', '%d')", author, playUrl, coverUrl, 0, 0, 0)
	db.Exec(Sql)
	defer db.Close()
}

// 获取视频信息
// get video inf
func GetVideo() []Video {
	var res []Video
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return res
	}
	rows, _ := db.Query("select * from video")
	var isFavorite bool
	var id, favoriteCount, commentCount int64
	var author, playUrl, coverUrl string
	for rows.Next() {
		rows.Scan(&id, &author, &playUrl, &coverUrl, &favoriteCount, &commentCount, &isFavorite)
		// fmt.Println(id, author, playUrl, coverUrl)
		userId, userName := GetIsExist(author)

		var user = User{
			Id:            userId,
			Name:          userName,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}

		var video = Video{
			Id:            id,
			Author:        user,
			PlayUrl:       playUrl, //构造完整视频链接。 create a whole video url.
			CoverUrl:      coverUrl,
			FavoriteCount: favoriteCount,
			CommentCount:  commentCount,
			IsFavorite:    false,
		}
		res = append([]Video{video}, res...) //视频倒叙存储
	}
	defer db.Close()
	// fmt.Println(len(res), "--------------------------------------------------")
	// if len(res) >= 30 { //视频长度大于30时，只取前30个
	// 	res = res[0:30]
	// }
	return res
}

//检测数据库中目标表是否存在，不存在则新建表， 存在则直接插入数据
func FavoriteTableChange(tableName string, userName string, action bool) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("create table if not exists %s (%s varchar(32))", tableName, "userName")
	db.Exec(Sql)
	if action {
		Sql = fmt.Sprintf("insert into %s values('%s')", tableName, userName)
	} else {
		Sql = fmt.Sprintf("delete from %s where userName='%s'", tableName, userName)
	}
	db.Exec(Sql)
	defer db.Close()
}

//通过视频ip查找视频喜爱列表
func GetVideoFavorite(videoId int64) []string {
	var res = []string{}
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return res
	}
	Sql := fmt.Sprintf("select * from favorite_%d", videoId)
	rows, err := db.Query(Sql)
	if err == nil {
		for rows.Next() {
			var temp string
			rows.Scan(&temp)
			// println(temp)
			res = append(res, temp)
		}
	}
	defer db.Close()
	// println(res)
	return res
}

//修改视频的喜爱数
func ChangeVideoFavorite(tableName string, rowName string, videoId int64, change string) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("update %s set %s=%s%s1 where videoId=%d", tableName, rowName, rowName, change, videoId)
	db.Exec(Sql)
	defer db.Close()
}

// 判断字符串是否在数组中
func StrINArr(str string, arr []string) bool {
	for _, i := range arr {
		if i == str {
			return true
		}
	}
	return false
}

// 获取评论信息
func GetComments(videoId int64) []Comment {
	var comments = []Comment{}
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return comments
	}
	Sql := `create table if not exists comments ( 
		commentId int unsigned NOT NULL AUTO_INCREMENT, 
		videoId int, author varchar(32), content varchar(255), 
		createDate varchar(20), PRIMARY KEY (commentId))`
	db.Exec(Sql)
	Sql = fmt.Sprintf("select * from comments where videoId='%d'", videoId)
	rows, err := db.Query(Sql)
	var commentId int64
	var author, content, createDate string
	if err == nil {
		for rows.Next() {
			rows.Scan(&commentId, &videoId, &author, &content, &createDate)
			userId, userName := GetIsExist(author)
			var user = User{
				Id:            userId,
				Name:          userName,
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false,
			}

			var comment = Comment{
				Id:         commentId,
				User:       user,
				Content:    content,
				CreateDate: createDate,
			}
			comments = append(comments, comment)
		}
	}
	defer db.Close()
	return comments
}

// 修改评论信息
func ChangeComment(commentId int64, videoId int64, author string, content string, createDate string, change bool) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	// Sql := "create table if not exists comments (commentId int unsigned NOT NULL AUTO_INCREMENT, videoId int, author varchar(32), content varchar(255), createDate varchar(20))"
	// db.Exec(Sql)
	var Sql string
	if change {
		Sql = fmt.Sprintf("insert into comments values (0, %d, '%s', '%s', '%s')", videoId, author, content, createDate)
	} else {
		Sql = fmt.Sprintf("delete from comments where videoId=%d and author='%s' and commentId='%d'", videoId, author, commentId)
	}
	db.Exec(Sql)
	defer db.Close()
}
