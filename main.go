package main

import (
	"github.com/RaymondCode/simple-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// x := test.GetAllFile("./publica")
	// for i := range x {
	// 	println(x[i])
	// }
}
