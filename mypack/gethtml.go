package mypack
import"github.com/gin-gonic/gin"
func Getinserthtml(c*gin.Context){
	c.HTML(200, "insert.html", nil)
}

func Getqueryhtml(c*gin.Context){
	c.HTML(200, "query.html", nil)
}
