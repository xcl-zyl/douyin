/*
点赞功能和用户个人信息获取点赞列表
favoriteAction and get FavoriteList in User personal interface
*/

package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	if _, exist := usersLoginInfo[token]; exist {
		// 根据当前红心状态，确认是取消还是点赞
		// 点赞成功添加视频ip至喜爱表中,视频表中喜爱数加1
		// 取消点赞即删除，减1
		video_id_int64, _ := strconv.ParseInt(video_id, 10, 64)
		if action_type == "1" {
			println("点赞")
			// 如果点赞已经点过，则跳过
			for _, value := range GetVideoFavorite(video_id_int64) {
				if value == usersLoginInfo[token].Name {
					c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User can't repeat operation"})
					return
				}
			}

			FavoriteTableChange("Favorite_"+video_id, token, true)
			ChangeVideoFavorite("video", "favoriteCount", video_id_int64, "+")
			ChangeUserFavorite("user", "favorite_count", token, "+")
			ChangeUserFavorite("user", "total_favorited", GetVideoAuthor(video_id_int64), "+")
		} else {
			println("取消")
			FavoriteTableChange("Favorite_"+video_id, token, false)
			ChangeVideoFavorite("video", "favoriteCount", video_id_int64, "-")
			ChangeUserFavorite("user", "favorite_count", token, "-")
			ChangeUserFavorite("user", "total_favorited", GetVideoAuthor(video_id_int64), "-")
		}
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
			// if usersLoginInfo[token].Name == value.Author.Name {
			// 	total_favorited += int(value.FavoriteCount)
			// }
		}
		// usersLoginInfo[token].Total_favorited = int64(total_favorited)

		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	token := c.Query("token")
	var videosOfUserFavorite = []Video{}
	for _, value := range DemoVideos {
		if StrINArr(usersLoginInfo[token].Name, GetVideoFavorite(value.Id)) {
			videosOfUserFavorite = append(videosOfUserFavorite, value)
		}
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videosOfUserFavorite,
	})
}
