package mypack

import (
	"github.com/bmizerany/assert"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mypack/mydatabase"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)
var r_v3 *gin.Engine
func init(){
	Init_maps()
	mydatabase.Init_db()
	mydatabase.Init_redis()
	achieve_v3:=Achieve_v3{}
	gin.DisableConsoleColor()
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r_v3= gin.Default()
	//长网址映射短网址
	r_v3.LoadHTMLGlob("C:/Users/Suncz/go/src/Shorturl/main/template/*")

	v1 :=r_v3.Group("")
	{
		v1.GET("/insert", func(c *gin.Context) {
			c.HTML(200, "insert.html", nil)
		})
		v1.POST("/insert", achieve_v3.Ltos)
		v1.GET("/redir/:url",achieve_v3.Stol)
		v1.GET("/query", func(c *gin.Context) {
			c.HTML(200, "query.html", nil)
		})
		//长网址查询短网址post请求
		v1.POST("/query", achieve_v3.Queryurl)
		//短网址查询长网址post请求
	}
}
func Get(url string, r *gin.Engine) []byte{
	//GET请求
	req := httptest.NewRequest("GET",url,nil)
	//初始化响应
	w:= httptest.NewRecorder()
	//调用相应接口
	r.ServeHTTP(w,req)
	res:=w.Result()
	defer res.Body.Close()
	//读取响应
	body,_:=ioutil.ReadAll(res.Body)
	return body
}
//构造POST
func Post(uri string, param url.Values, r *gin.Engine) []byte {
	// 构造post请求
	req := httptest.NewRequest("POST", uri, strings.NewReader(param.Encode()))
	//req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//fmt.Println(req)
	// 初始化响应
	w := httptest.NewRecorder()
	// 调用相应handler接口
	r.ServeHTTP(w, req)

	// 提取响应
	result := w.Result()

	defer result.Body.Close()

	// 读取响应body
	body, _ := ioutil.ReadAll(result.Body)
	//fmt.Println(body)
	return body
}

func TestAchieve_v3_Ltos(t *testing.T) {
	uri := "/insert"
	param := url.Values{
		"longurl":{"www.baidu.com"},
	}
	//fmt.Println(param)
	var str string = "{\"formaterror\":\"need http:// or https://\"}"
	body := Post(uri,param,r_v3)
	test_str := string(body[:])
	assert.Equal(t,str,test_str)
	str = "{\"shorturl\":\"existed\"}"
	param = url.Values{
		"longurl":{"http://www.baidu.com"},
	}
	body = Post(uri,param,r_v3)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
}

func TestAchieve_v3_Stol(t *testing.T) {
	url_2 := "/redir/2c"
	str := "{\"error\":\"not found\"}"
	body :=Get(url_2,r_v3)
	test_str := string(body[:])
	assert.Equal(t,str,test_str)
}

func TestAchieve_v3_Queryurl(t *testing.T) {
	uri := "/query"
	param := url.Values{
		"longurl":{"www.baidu.com"},
	}
	var str string = "{\"formaterror\":\"need http:// or https://\"}"
	body := Post(uri,param,r_v3)
	test_str := string(body[:])
	assert.Equal(t,str,test_str)
	str = "{\"shorturl\":\"1\"}"
	param = url.Values{
		"longurl":{"http://www.baidu.com"},
	}
	body = Post(uri,param,r_v3)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
	param = url.Values{
		"shorturl":{"wwww."},
	}
	str = "{\"error\":\"not found\"}"
	body = Post(uri,param,r_v3)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
	str = "{\"longurl\":\"http://www.baidu.com\"}"
	param = url.Values{
		"shorturl":{"1"},
	}
	body =Post(uri,param,r_v3)
	test_str = string(body[:])
	assert.Equal(t,str,test_str)
}

func TestGetinserthtml(t *testing.T) {
	uri := "/insert"

	// 发起Get请求
	body := Get(uri, r_v3)
	//fmt.Printf("response:%v\n", string(body))

	// 判断响应是否与预期一致
	html,_ := ioutil.ReadFile("C:/Users/Suncz/Goproject/src/pack/template/insert.html")
	htmlStr := string(html)

	if htmlStr != string(body) {
		t.Errorf("响应数据不符，body:%v\n",string(body))
	}
}

func TestGetqueryhtml(t *testing.T) {
	uri := "/query"

	// 发起Get请求
	body := Get(uri, r_v3)
	//fmt.Printf("response:%v\n", string(body))
	// 判断响应是否与预期一致
	html, _ := ioutil.ReadFile("C:/Users/Suncz/Goproject/src/pack/template/query.html")
	htmlStr := string(html)
    assert.Equal(t,string(body),htmlStr)
	if htmlStr != string(body) {
		t.Errorf("响应数据不符，body:%v\n", string(body))
	}
}

