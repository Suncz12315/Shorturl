package mypack

import (
"github.com/bmizerany/assert"
"github.com/gin-gonic/gin"
"net/url"
"testing"
"mypack/mydatabase"
"io"
"os"
)
var r_v2 *gin.Engine
func init(){
	Init_maps()
	mydatabase.Init_db()
	achieve_v2:=Achieve_v2{}
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r_v2= gin.Default()
	//长网址映射短网址
	r_v2.LoadHTMLGlob("C:/Users/Suncz/go/src/Shorturl/main/template/*")

	v1 :=r_v2.Group("")
	{
		v1.GET("/insert", func(c *gin.Context) {
			c.HTML(200, "insert.html", nil)
		})
		v1.POST("/insert", achieve_v2.Ltos)
		v1.GET("/redir/:url",achieve_v2.Stol)
		v1.GET("/query", func(c *gin.Context) {
			c.HTML(200, "query.html", nil)
		})
		//长网址查询短网址post请求
		v1.POST("/query", achieve_v2.Queryurl)
		//短网址查询长网址post请求
	}
}

func TestAchieve_v2_Ltos(t *testing.T) {
	uri := "/insert"
	param := url.Values{
		"longurl":{"www.baidu.com"},
	}
	//fmt.Println(param)
	var str string = "{\"formaterror\":\"need http:// or https://\"}"
	body := Post(uri,param,r_v2)
	test_str := string(body[:])
	assert.Equal(t,str,test_str)
	str = "{\"shorturl\":\"existed\"}"
	param = url.Values{
		"longurl":{"http://www.baidu.com"},
	}
	body = Post(uri,param,r_v2)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
}

func TestAchieve_v2_Stol(t *testing.T) {
	url_2 := "/redir/2c"
	str := "{\"error\":\"not found\"}"
	body :=Get(url_2,r_v2)
	test_str := string(body[:])
	assert.Equal(t,str,test_str)
}

func TestAchieve_v2_Queryurl(t *testing.T) {
	uri := "/query"
	param := url.Values{
		"longurl":{"www.baidu.com"},
	}
	var str string = "{\"formaterror\":\"need http:// or https://\"}"
	body := Post(uri,param,r_v2)
	test_str := string(body[:])
	assert.Equal(t,str,test_str)
	str = "{\"shorturl\":\"1\"}"
	param = url.Values{
		"longurl":{"http://www.baidu.com"},
	}
	body = Post(uri,param,r_v2)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
	param = url.Values{
		"shorturl":{"wwww."},
	}
	str = "{\"error\":\"not found\"}"
	body = Post(uri,param,r_v2)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
	str = "{\"longurl\":\"http://www.baidu.com\"}"
	param = url.Values{
		"shorturl":{"1"},
	}
	body =Post(uri,param,r_v2)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
}
