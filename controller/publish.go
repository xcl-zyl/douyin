/*
视频发布功能和获取用户发布视频列表
publish videos and get all published videos of user
*/

package controller

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)

	var video = Video{
		Id:            int64(len(DemoVideos) + 1),
		Author:        user,
		PlayUrl:       HostIp + finalName, //构造完整视频链接。 create a whole video url.
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	}
	DemoVideos = append([]Video{video}, DemoVideos...)         //新发布的视频放在开头
	AddVideo(video.Author.Name, video.PlayUrl, video.CoverUrl) // 将发布视频加入数据库中

	saveFile := filepath.Join("./public/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	// println(token)
	var videosOfUser = []Video{}
	// println(len(usersLoginInfo))
	for _, value := range DemoVideos {
		if value.Author.Name == usersLoginInfo[token].Name {
			// fmt.Println(value.Author.Name)
			// fmt.Println(usersLoginInfo[token].Name)
			// fmt.Println(value)
			videosOfUser = append(videosOfUser, value)
		}
	}
	// println(len(videosOfUser))
	// if len(videosOfUser) == 0 {
	// 	videosOfUser = []Video{}
	// }
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videosOfUser,
	})
}
