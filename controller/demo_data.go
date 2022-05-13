/*
数据文件：包含host ip信息， 所有视频信息
data: include host ip, all videos information
*/

package controller

// 原本视频列表是写死的，这里改为根据文件夹内容动态改变
// At first, the video array is changeless, I make them change with the video folder.
// var DemoVideos = []Video{
// 	{
// 		Id:            1,
// 		Author:        DemoUser,
// 		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
// 		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
// 		FavoriteCount: 0,
// 		CommentCount:  0,
// 		IsFavorite:    false,
// 	},
// 	// {
// 	// 	Id:            2,
// 	// 	Author:        DemoUser,
// 	// 	PlayUrl:       hostIp,
// 	// 	CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
// 	// 	FavoriteCount: 0,
// 	// 	CommentCount:  0,
// 	// 	IsFavorite:    false,
// 	// },
// }

// 先获取主机ip， 构造视频文件夹链接
// First, we need to get the host ip using func GetHostIp in the tool file， to create the video folder url.
var HostIp string = func() string {
	ip := "http://" + GetHostIp() + ":8080/static/"
	// println(ip + "                     ---------------------------------")
	return ip
}()

// 通过public文件夹，动态生成视频列表
// create DemoVideos using files in pubilc folder.
// var DemoVideos = func() []Video {
// 	var demoVideos = []Video{}
// 	videoName := GetAllFile("./public")
// 	for i := range videoName {
// 		var video = Video{
// 			Id:            int64(i + 1),
// 			Author:        DemoUser,
// 			PlayUrl:       HostIp + videoName[i], //构造完整视频链接。 create a whole video url.
// 			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
// 			FavoriteCount: 0,
// 			CommentCount:  0,
// 			IsFavorite:    false,
// 		}
// 		demoVideos = append(demoVideos, video)
// 	}
// 	return demoVideos
// }()

var DemoVideos = GetVideo()

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "test",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
