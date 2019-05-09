package mypack
import(
	"github.com/gin-gonic/gin"
)
type Achieve_v1 struct{
    
}

func (this *Achieve_v1)Ltos(c*gin.Context){
	var form InsertForm
	if c.ShouldBind(&form) !=nil{
		panic("bind error")
	}
	url :=form.Longurl
	if (Isvalid_url(url)) {
		lk.Lock()
		if (stoi_maps[url] == 0) {
			cnt++
			stoi_maps[url] = cnt
			itos_maps[cnt] = url
		}
		lk.Unlock()

		var s string = To62mode(stoi_maps[url])
		c.JSON(200, gin.H{
			"shorturl": s,
		})
	} else {
		c.JSON(200, gin.H{
			"formaterror": "need http:// or https://",
		})
	}
}

func (this *Achieve_v1)Stol(c*gin.Context){
	url :=c.Param("url")
	num:= To10mode(url) //找到10进制的数
	if (num != 0 ){
		if(itos_maps[num]!="") {
			c.Redirect(301, itos_maps[num])
		}else {
			c.JSON(200, gin.H{
				"error":"not found",
			})
		}
	} else{
		c.JSON(200,gin.H{
			"error":"formaterror",
		})
	}
}
func (this *Achieve_v1)Queryurl(c*gin.Context){
	//url :=c.PostForm("url")
	var form QueryForm
	if c.ShouldBind(&form)!=nil{
		panic("bind error")
	}
	//c.ShouldBind(&form)
	if form.Longurl!=""&&form.Shorturl=="" {
		long_url := form.Longurl
		if Isvalid_url(long_url) {
			num := stoi_maps[long_url]
			if (num != 0) {
				shorturl := To62mode(num)
				c.JSON(200, gin.H{
					"shorturl": shorturl,
				})
			} else {
				//在没有对应断短网址时，考虑添加？
				c.JSON(200, gin.H{
					"error": "not found",
				})
			}
		} else {
			c.JSON(200, gin.H{
				"formaterror": "need http:// or https://",
			})
		}
	}else if form.Shorturl!=""&&form.Longurl=="" {
		short_url := form.Shorturl
		num := To10mode(short_url) //找到10进制的数
		if (num == 0) {
			c.JSON(200, gin.H{
				"error": "formaterror", //短网址包含非法输入
			})
		} else if (itos_maps[num] != "") {
			c.JSON(200, gin.H{
				"longurl": itos_maps[num],
			})
		} else {
			c.JSON(200, gin.H{
				"error": "not found",
			})
		}
	}else{
		c.JSON(200,gin.H{
			"error":"input error",
		})
	}
}