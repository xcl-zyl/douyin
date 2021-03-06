/*
评论：包含发送评论，获取评论列表
comment： include send comment, get all comments in the video
*/

package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	Response
	Comment `json:"comment,omitempty"`
}

var Comments = []Comment{}

// CommentAction no practical effect, just check if token is valid
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")
	comment_text := c.Query("comment_text")
	comment_id := c.Query("comment_id")

	if _, exist := usersLoginInfo[token]; exist {
		video_id_int64, _ := strconv.ParseInt(video_id, 10, 64)
		comment_id_int64, _ := strconv.ParseInt(comment_id, 10, 64)
		createDate := time.Now().Format("01-02")
		if action_type == "1" {
			println("评论")
			comment_id_int64 = ChangeComment(comment_id_int64, video_id_int64, usersLoginInfo[token].Name, comment_text, createDate, true)
			ChangeVideoFavorite("video", "commentCount", video_id_int64, "+")
			var comment = Comment{
				Id:         comment_id_int64,
				User:       usersLoginInfo[token],
				Content:    comment_text,
				CreateDate: createDate,
			}
			c.JSON(http.StatusOK, CommentResponse{
				Response: Response{StatusCode: 0},
				Comment:  comment,
			})
		} else {
			println("删除")
			ChangeComment(comment_id_int64, video_id_int64, usersLoginInfo[token].Name, comment_text, createDate, false)
			ChangeVideoFavorite("video", "commentCount", video_id_int64, "-")
			c.JSON(http.StatusOK, Response{StatusCode: 0})
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
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
func CommentList(c *gin.Context) {
	video_id := c.Query("video_id")
	video_id_int64, _ := strconv.ParseInt(video_id, 10, 64)
	Comments = GetComments(video_id_int64)
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: Comments,
	})
}
