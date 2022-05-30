/*
关注功能， 获取关注列表和获取粉丝列表
RelationAction， FollowList， FollowerList
*/

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	to_user_id := c.Query(("to_user_id"))
	action_type := c.Query("action_type")

	if _, exist := usersLoginInfo[token]; exist {
		to_user_id_int64, _ := strconv.ParseInt(to_user_id, 10, 64)
		if usersLoginInfo[token].Id == to_user_id_int64 {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User can't focus on yourself"})
			return
		}
		if action_type == "1" {
			println("关注")
			// 如果关注对象已经关注，则跳过
			for _, j := range GetUserFollowAndFollower(usersLoginInfo[token].Name, "follow") {
				if GetUserName(to_user_id_int64) == j.Name {
					c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User can't repeat operation"})
					return
				}
			}

			ChangeFollowAndFollower(token, to_user_id_int64, true)
			ChangeUserFollowAndFollowerNum(token, to_user_id_int64, "+")
		} else {
			println("取关")
			ChangeFollowAndFollower(token, to_user_id_int64, false)
			ChangeUserFollowAndFollowerNum(token, to_user_id_int64, "-")
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
		DemoVideos = GetVideo()
		for i, value := range DemoVideos {
			if token != "" && StrINArr(usersLoginInfo[token].Name, GetVideoFavorite(value.Id)) {
				DemoVideos[i].IsFavorite = true
			}
			for _, j := range GetUserFollowAndFollower(usersLoginInfo[token].Name, "follow") {
				if j.Name == value.Author.Name {
					DemoVideos[i].Author.IsFollow = true
					break
				}
			}
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// UserList: []User{usersLoginInfo[token]},
		UserList: GetUserFollowAndFollower(usersLoginInfo[token].Name, "follow"),
	})
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	followers := GetUserFollowAndFollower(usersLoginInfo[token].Name, "follower")
	for i, value := range followers {
		for _, j := range GetUserFollowAndFollower(usersLoginInfo[token].Name, "follow") {
			if value.Name == j.Name {
				followers[i].IsFollow = true
				break
			}
		}
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		// UserList: []User{usersLoginInfo[token]},
		UserList: followers,
	})
}
