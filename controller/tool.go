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
// 补充数据库添加新的内容
func GetIsExist(username string, password ...string) (int64, string, int64, int64, int64, int64) {
	var userId, followCount, followerCount, favorite_count, total_favorited int64
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败" + err.Error())
		return userId, username, followCount, followerCount, favorite_count, total_favorited
	}

	if len(password) != 0 { // select 查了多少元素 scan 就必须读多少元素， 个数必须统一
		Sql := fmt.Sprintf("select * from user where userName = '%s' and password = '%s'", username, password[0])
		db.QueryRow(Sql).Scan(&userId, &username, &password[0], &followCount, &followerCount, &favorite_count, &total_favorited)
		// fmt.Println(userId, username, password, followCount, followerCount)
	} else {
		var pass string
		Sql := fmt.Sprintf("select * from user where userName = '%s'", username)
		db.QueryRow(Sql).Scan(&userId, &username, &pass, &followCount, &followerCount, &favorite_count, &total_favorited)
	}
	defer db.Close()
	return userId, username, followCount, followerCount, favorite_count, total_favorited
}

// 通过id查用户名
func GetUserName(userId int64) string {
	var userName string
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败" + err.Error())
		return userName
	}
	Sql := fmt.Sprintf("select userId, userName from user where userId=%d", userId)
	db.QueryRow(Sql).Scan(&userId, &userName)
	defer db.Close()
	return userName
}

// 向数据库插入用户信息
// insert user in mysql
func AddUser(username, password string) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("insert into user values (0, '%s', '%s', 0, 0, 0, 0)", username, password)
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

func GetVideoAuthor(id int64) string {
	var author string
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return author
	}
	Sql := fmt.Sprintf("select videoId, author from video where videoId=%d", id)
	db.QueryRow(Sql).Scan(&id, &author)
	defer db.Close()
	return author
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
	var id, favoriteCount, commentCount int64
	var author, playUrl, coverUrl string
	for rows.Next() {
		rows.Scan(&id, &author, &playUrl, &coverUrl, &favoriteCount, &commentCount)
		// fmt.Println(id, author, playUrl, coverUrl)
		userId, userName, followCount, followerCount, favorite_count, total_favorited := GetIsExist(author)

		var user = User{
			Id:              userId,
			Name:            userName,
			FollowCount:     followCount,
			FollowerCount:   followerCount,
			Favorite_count:  favorite_count,
			Total_favorited: total_favorited,
			IsFollow:        false,
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

// 修改视频的喜爱数
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

// 修改用户喜爱数
func ChangeUserFavorite(tableName string, rowName string, userName string, change string) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("update %s set %s=%s%s1 where userName='%s'", tableName, rowName, rowName, change, userName)
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
			userId, userName, followCount, follower_count, favorite_count, total_favorited := GetIsExist(author)
			var user = User{
				Id:              userId,
				Name:            userName,
				FollowCount:     followCount,
				FollowerCount:   follower_count,
				Favorite_count:  favorite_count,
				Total_favorited: total_favorited,
				IsFollow:        false,
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
func ChangeComment(commentId int64, videoId int64, author string, content string, createDate string, change bool) int64 {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return 0
	}
	// Sql := "create table if not exists comments (commentId int unsigned NOT NULL AUTO_INCREMENT, videoId int, author varchar(32), content varchar(255), createDate varchar(20))"
	// db.Exec(Sql)
	var Sql string
	if change {
		Sql = fmt.Sprintf("insert into comments values (0, %d, '%s', '%s', '%s')", videoId, author, content, createDate)
	} else {
		Sql = fmt.Sprintf("delete from comments where videoId=%d and author='%s' and commentId='%d'", videoId, author, commentId)
	}
	res, _ := db.Exec(Sql)
	commentId, _ = res.LastInsertId()
	defer db.Close()
	return commentId
}

// 修改用户粉丝表以及关注表
func ChangeFollowAndFollower(userName string, to_user_id int64, change bool) {
	to_user_name := GetUserName(to_user_id)
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("create table if not exists follow_%s (%s varchar(32))", userName, "followName")
	db.Exec(Sql)
	Sql = fmt.Sprintf("create table if not exists follower_%s (%s varchar(32))", to_user_name, "followerName")
	db.Exec(Sql)

	var Sql1, Sql2 string
	if change {
		Sql1 = fmt.Sprintf("insert into follow_%s values ('%s')", userName, to_user_name)
		Sql2 = fmt.Sprintf("insert into follower_%s values ('%s')", to_user_name, userName)
	} else {
		Sql1 = fmt.Sprintf("delete from follow_%s where followName='%s'", userName, to_user_name)
		Sql2 = fmt.Sprintf("delete from follower_%s where followerName='%s'", to_user_name, userName)
	}
	db.Exec(Sql1)
	db.Exec(Sql2)

	defer db.Close()
}

// 修改用户粉丝数以及关注数
func ChangeUserFollowAndFollowerNum(userName string, to_user_id int64, change string) {
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}
	Sql := fmt.Sprintf("update %s set %s=%s%s1 where userName='%s'", "user", "followCount", "followCount", change, userName)
	db.Exec(Sql)
	Sql = fmt.Sprintf("update %s set %s=%s%s1 where userId=%d", "user", "followerCount", "followerCount", change, to_user_id)
	db.Exec(Sql)
	defer db.Close()
}

//读取用户粉丝或者关注者列表
func GetUserFollowAndFollower(userName string, object string) []User {
	// user_id, _, _, _ := GetIsExist(userName)
	var objects = []User{}
	db, err := sql.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println("数据库连接失败")
		return objects
	}
	var Sql string
	var isFollow bool
	if object == "follow" {
		Sql = fmt.Sprintf("create table if not exists follow_%s (%s varchar(32))", userName, "followName")
		db.Exec(Sql)
		Sql = fmt.Sprintf("select * from follow_%s", userName)
		isFollow = true
	} else {
		Sql = fmt.Sprintf("create table if not exists follower_%s (%s varchar(32))", userName, "followerName")
		db.Exec(Sql)
		Sql = fmt.Sprintf("select * from follower_%s", userName)
		isFollow = false
	}
	var objectName string
	rows, err := db.Query(Sql)
	if err == nil {
		for rows.Next() {
			rows.Scan(&objectName)

			userId, userName, followCount, follower_count, favorite_count, total_favorited := GetIsExist(objectName)
			var user = User{
				Id:              userId,
				Name:            userName,
				FollowCount:     followCount,
				FollowerCount:   follower_count,
				Favorite_count:  favorite_count,
				Total_favorited: total_favorited,
				IsFollow:        isFollow,
			}
			objects = append(objects, user)
		}
	}

	defer db.Close()
	return objects
}
