package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	// 请求类型
	r.GET("/gobasic", func(context *gin.Context) {
		context.String(http.StatusOK, "method:get")
	})
	r.POST("/gobasic", func(context *gin.Context) {
		context.String(http.StatusOK, "method:post")
	})
	r.PUT("/gobasic", func(context *gin.Context) {
		context.String(http.StatusOK, "method:put")
	})
	r.DELETE("/gobasic", func(context *gin.Context) {
		context.String(http.StatusOK, "method:delete")
	})

	// 都是前缀匹配
	r.GET("/user", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello, Gocobnus")
	})

	r.GET("/user/:name/*role", func(context *gin.Context) {
		name := context.Param("name") // 获取url中的参数值信息
		context.String(http.StatusOK, "Hello, %s", name)
	})

	// curl -X GET http://localhost:8080/userinfo\?name\=gocobnus\&age\=25
	r.GET("/userinfo", func(context *gin.Context) {
		name := context.Query("name")
		age := context.DefaultQuery("age", "18")
		context.String(http.StatusOK, "user info: name:%s, age:%s", name, age)
	})

	// POST
	// curl -X POST http://localhost:8080/form_post -d 'username=gocobnus&password=123456'
	r.POST("form_post", func(context *gin.Context) {
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "000000")
		context.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
	// POST&GET
	// curl -X POST http://localhost:8080/mix_post_get\?id\=10086 -d 'username=gocobnus&password=123456'
	r.POST("mix_post_get", func(context *gin.Context) {
		id := context.Query("id")
		username := context.PostForm("username")
		password := context.DefaultPostForm("password", "000000")
		context.JSON(http.StatusOK, gin.H{
			"id":       id,
			"username": username,
			"password": password,
		})
	})

	// 默认绑定8080
	r.Run()

}
