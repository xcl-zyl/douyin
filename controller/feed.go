/*
打开软件获取视频流
open program to get video data
*/

package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	demoVideos := DemoVideos //获取视频流时检查视频数量
	if len(demoVideos) >= 30 {
		demoVideos = demoVideos[0:30]
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: demoVideos,
		NextTime:  time.Now().Unix(),
	})
}
