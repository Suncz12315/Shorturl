package mypack
import (
	"github.com/gin-gonic/gin"
	"mypack/mydatabase"
	"reflect"
)
type Achieve_v2 struct{

}
//插入长网址
func (this Achieve_v2)Ltos(c*gin.Context){
	//url := c.PostForm("url")
	var url InsertForm
	if c.ShouldBind(&url)!=nil{
		panic("bind error")
	}
	//判断是否为合法网址
	if !Isvalid_url(url.Longurl){
		c.JSON(200, gin.H{
			"formaterror": "need http:// or https://",
		})
		return 
	}
	//判断url长度是否过长
	if (len(url.Longurl)>9216){
		c.JSON(200,gin.H{
			"error": "the length of URL is too long",
		})
		return
	}
	lk.Lock()
	last_search_res := mydatabase.Db.Model(&mydatabase.Ltostable{}).Last(&mydatabase.Ltostable{})
	res, ok := last_search_res.Value.(*mydatabase.Ltostable)
	if ok==false{
		panic("last query error")
	}
	shorturl := To62mode(res.URL_id + 1)
	item := mydatabase.Ltostable{Longurl: url.Longurl, Shorturl: shorturl}
	cs := mydatabase.Db.Create(&item)
	if reflect.TypeOf(cs.Error) != nil {
		c.JSON(200, gin.H{
			"shorturl": "existed",
		})
	} else {
		c.JSON(200, gin.H{
			"shorturl": shorturl,
		})
	}
	lk.Unlock()
}
//重定向
func (this Achieve_v2)Stol(c*gin.Context){
	url :=c.Param("url")
	id := mydatabase.Db.Model(&mydatabase.Ltostable{}).Where("shorturl=?", url).Find(&mydatabase.Ltostable{})
	v, ok := id.Value.(*mydatabase.Ltostable)
	if ok ==false{
		panic("query error")
	}
	if v.Shorturl == "" {
		c.JSON(200, gin.H{
			"error": "not found",
		})
	} else {
		c.Redirect(302, v.Longurl)
	}
}
//查询
func (this Achieve_v2)Queryurl(c*gin.Context){
	var form QueryForm
	if c.ShouldBind(&form)!=nil{
		panic("bind error")
	}
	if form.Longurl!=""&&form.Shorturl==""{
		//输入长网址是否合法
		if !Isvalid_url(form.Longurl){
			c.JSON(200, gin.H{
				"formaterror": "need http:// or https://",
			})
			return
		}
		search_res := mydatabase.Db.Model(&mydatabase.Ltostable{}).Where("longurl=?", form.Longurl).Find(&mydatabase.Ltostable{})
		res, ok := search_res.Value.(*mydatabase.Ltostable)
		if ok==false{
			panic("query error")
		}
		if res.Longurl == "" {
			c.JSON(200, gin.H{
				"error": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"shorturl": res.Shorturl,
			})
		}
	}else if form.Shorturl!=""&&form.Longurl==""{
		search_res := mydatabase.Db.Model(&mydatabase.Ltostable{}).Where("shorturl=?", form.Shorturl).Find(&mydatabase.Ltostable{})
		res, ok := search_res.Value.(*mydatabase.Ltostable)
		if ok == false{
			panic("query error")
		}
		if res.Shorturl == "" {
			c.JSON(200, gin.H{
				"error": "not found",
			})
		} else {
			c.JSON(200, gin.H{
				"longurl": res.Longurl,
			})
		}
	}else{
		c.JSON(200,gin.H{
			"error":"input error",
		})
	}
}