/*
登录注册以及获取用户信息
*/

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
// after add db, need to read db
var usersLoginInfo = map[string]User{}

// var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	// token := username + password
	if exist, _ := GetIsExist(username); exist != 0 {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		// atomic.AddInt64(&userIdSequence, 1)
		// newUser := User{
		// 	Id:   userIdSequence,
		// 	Name: username,
		// }
		// usersLoginInfo[token] = newUser

		//token 只需要用户名，安全性较低，待优化
		AddUser(username, password)
		userId, _ := GetIsExist(username)
		user := User{
			Id:            userId,
			Name:          username,
			FollowCount:   0,
			FollowerCount: 0,
			IsFollow:      false,
		}
		usersLoginInfo[username] = user
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userId,
			Token:    username,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if exist, _ := GetIsExist(username); exist != 0 {
		if exist, _ := GetIsExist(username, password); exist != 0 {
			favorite_count := 0
			total_favorited := 0
			for i, value := range DemoVideos {
				if StrINArr(username, GetVideoFavorite(value.Id)) {
					DemoVideos[i].IsFavorite = true
					favorite_count++
				}
				if username == value.Author.Name {
					total_favorited += int(value.FavoriteCount)
				}
			}
			user := User{
				Id:              exist,
				Name:            username,
				FollowCount:     0,
				FollowerCount:   0,
				Favorite_count:  int64(favorite_count),
				Total_favorited: int64(total_favorited),
				IsFollow:        false,
			}
			usersLoginInfo[username] = user
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   exist,
				Token:    username,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Password is wrong"},
			})
		}

	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	total_favorited := 0
	for _, value := range DemoVideos {
		if usersLoginInfo[token].Name == value.Author.Name {
			total_favorited += int(value.FavoriteCount)
		}
	}
	user := usersLoginInfo[token]
	user.Total_favorited = int64(total_favorited)
	// usersLoginInfo[token].Total_favorited = int64(total_favorited)
	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
