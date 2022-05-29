/*
启动文件
main file, start program

router: 路由表，包含全部的api调用接口
router: include all api interface
gin框架：Go的高效HTTP Web框架，具体使用方法打开"github.com/gin-gonic/gin"即可看到完整说明文档
gin: a efficient HTTP Web frame, you can open "github.com/gin-gonic/gin" to know how to use it.
*/

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xcl-zyl/douyin/router"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// controller.GetComments(7)
	// controller.AddComment(7, "aab", "aaa", "05-01")
	// for _, i := range controller.GetComments(7) {
	// 	fmt.Println(i)
	// }
	// controller.ChangeComment(7, "aab", "aaa", "05-01", false)
	// for _, i := range controller.GetComments(7) {
	// 	fmt.Println(i)
	// }
	// controller.ChangeComment(0, 7, "aaa", "bbb", "05-01", true)
	// for _, i := range controller.GetComments(7) {
	// 	fmt.Println(i)
	// }
	// controller.ChangeComment(1, 7, "aaa", "bbb", "05-01", false)
	// for _, i := range controller.GetComments(7) {
	// 	fmt.Println(i)
	// }
	// test.GetVideoFavorite(5)
	// test.TableIsExist("bbb", 3)
	// test.FavoriteTableChange("aaa", "bbb", false)
	//test.AddUser("abc", "123")
	// println(test.GetIsExist("abc"))
	// println(test.GetIsExist("ab"))
	// println(test.GetIsExist("abc", "123456"))
	// println(test.GetIsExist("abc", "12345"))
	// println(test.GetIsExist("zhanglei", "douyin"))
	// test.TestSlice()
}
