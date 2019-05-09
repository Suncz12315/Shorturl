package mypack
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mypack/mydatabase"
	"reflect"
)
type Achieve_v3 struct{

}
func (this Achieve_v3)Ltos(c*gin.Context){
	var url InsertForm
	if c.ShouldBind(&url)!=nil{
		panic("bind error")
	}
	if !Isvalid_url(url.Longurl){
		c.JSON(200, gin.H{
			"formaterror": "need http:// or https://",
		})
		return
	}
	item,_ := mydatabase.Cli_redis.Get(url.Longurl).Result() //在redis缓存中查询提交的长网址是否存在
	if item !=""{
		c.JSON(200, gin.H{
			"shorturl": "existed",
		})
		return
	}
	lk.Lock()
	get_value := mydatabase.Cli_redis.Get("count_url")
	value,_ := get_value.Result() //获取计数器的值
	var tmp_count int64
	if value == ""|| value == "0"{ //count_url不在redis中，找数据库中最后的id获取
	    search_res := mydatabase.Db.Model(&mydatabase.Ltostable{}).Last(&mydatabase.Ltostable{})
	    search_count,ok := search_res.Value.(*mydatabase.Ltostable)
	    if ok==false{
	    	panic("query error")
		}
		tmp_count = search_count.URL_id
		fmt.Printf("insert count_url success")
		mydatabase.Cli_redis.Set("count_url",tmp_count,0)
	}else{
		tmp_count,_ = get_value.Int64()
	}
	short_url := To62mode(tmp_count+1)                                         //转为62进制
	db_item := mydatabase.Ltostable{Longurl: url.Longurl, Shorturl: short_url} //数据库数据
	cs := mydatabase.Db.Create(&db_item)
	if reflect.TypeOf(cs.Error) != nil { //在mysql数据库中已存在
	    c.JSON(200, gin.H{
	    	"shorturl": "existed",
	    })
	    mydatabase.Cli_redis.Set(url.Longurl,short_url,0) //插入已存在于数据库中但不在redis中的数据，因此判定为热点数据放入redis
	    } else{
	    	long_insert_err := mydatabase.Cli_redis.Set(url.Longurl,short_url,0).Err() //redis中长网址-》短网址
	    	if long_insert_err!=nil {
	    		panic(long_insert_err)
	    		}
	    	short_insert_err:= mydatabase.Cli_redis.Set(short_url,url.Longurl,0).Err() //redis中短网址-》长网址
	    	if short_insert_err!=nil{
	    		panic(short_insert_err)
	    	}
	    	mydatabase.Cli_redis.Incr("count_url") //计数器自增
	    	c.JSON(200, gin.H{
	    		"shorturl": short_url,
	    	})
	    }
	lk.Unlock()
}

func (this Achieve_v3)Stol(c*gin.Context){
	url := c.Param("url")
	item, _ := mydatabase.Cli_redis.Get(url).Result()
	if item == ""{
		id := mydatabase.Db.Model(&mydatabase.Ltostable{}).Where("shorturl=?", url).Find(&mydatabase.Ltostable{})
		v, ok := id.Value.(*mydatabase.Ltostable)
		if ok ==false {
			panic("query error")
		}
		if v.Shorturl == "" {
			c.JSON(200, gin.H{
				"error": "not found",
			})
		} else {
			mydatabase.Cli_redis.Set(url,v.Longurl,0) //将此次在数据库中找到的重定向结果放入redis
			c.Redirect(302, v.Longurl)
		}
	}else{
		c.Redirect(302, item)
	}
}

func (this Achieve_v3)Queryurl(c*gin.Context){
	var form QueryForm
	if c.ShouldBind(&form)!=nil{
		panic("bind error")
	}
	if form.Longurl != ""&&form.Shorturl=="" {
		if !Isvalid_url(form.Longurl) {
			c.JSON(200, gin.H{
				"formaterror": "need http:// or https://",
			})
			return
		}
		short_url, _ := mydatabase.Cli_redis.Get(form.Longurl).Result()
		if short_url == "" {
			search_res := mydatabase.Db.Model(&mydatabase.Ltostable{}).Where("longurl=?", form.Longurl).Find(&mydatabase.Ltostable{})
			res, ok := search_res.Value.(*mydatabase.Ltostable)
			if ok == false {
				panic("query error")
			}
			if res.Longurl == "" {
				c.JSON(200, gin.H{
					"error": "not found",
				})
			} else {
				mydatabase.Cli_redis.Set(form.Longurl, res.Shorturl, 0)
				c.JSON(200, gin.H{
					"shorturl": res.Shorturl,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"shorturl": short_url,
			})
		}
	} else if form.Shorturl != "" &&form.Longurl == ""{
			long_url, _ := mydatabase.Cli_redis.Get(form.Shorturl).Result()
			if long_url == "" {
				search_res := mydatabase.Db.Model(&mydatabase.Ltostable{}).Where("shorturl=?", form.Shorturl).Find(&mydatabase.Ltostable{})
				res, ok := search_res.Value.(*mydatabase.Ltostable)
				if ok==false{
					panic("query error")
				}
				if res.Shorturl == "" {
					c.JSON(200, gin.H{
						"error": "not found",
					})
				} else {
					mydatabase.Cli_redis.Set(form.Shorturl, res.Longurl, 0)
					c.JSON(200, gin.H{
						"longurl": res.Longurl,
					})
				}
			} else {
				c.JSON(200, gin.H{
					"longurl": long_url,
				})
			}
		} else {
			c.JSON(200, gin.H{
				"error": "invalid input",
			})
		}
}

