package mypack

import (
	"github.com/gin-gonic/gin"
	"io"
	"mypack/mydatabase"
	"os"
)
//路由设计
func Web_init(achieve_interface Achieve_interface){
	gin.DisableConsoleColor()

	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.LoadHTMLGlob("C:/Users/Suncz/go/src/Shorturl/main/template/*")

	v1 :=r.Group("")
	{
		/*v1.GET("/insert", func(c *gin.Context) {
			c.HTML(200, "insert.html", nil)
		})*/
		v1.GET("/insert",Getinserthtml)
		v1.POST("/insert", achieve_interface.Ltos)
		v1.GET("/redir/:url",achieve_interface.Stol)
		/*v1.GET("/query", func(c *gin.Context) {
			c.HTML(200, "query.html", nil)
		})*/
		v1.GET("/query",Getqueryhtml)
		v1.POST("/query", achieve_interface.Queryurl)
	}
	r.Run(":8080")
}
//跑第一个版本
func Run_v1(){
	Init_maps()
	achieve_v1:=Achieve_v1{}
	Web_init(&achieve_v1)
}


//跑第二版本
func Run_v2(){
	mydatabase.Init_db()
	Init_maps()
	achieve_v2 := Achieve_v2{}
    Web_init(&achieve_v2)
    mydatabase.Close_db()
}

func Run_v3(){
	Init_maps()
	mydatabase.Init_redis()
	mydatabase.Init_db()
	achieve_v3 := Achieve_v3{}
	Web_init(&achieve_v3)
	mydatabase.Close_db()
}


