package mypack
import"github.com/gin-gonic/gin"
//定义实现接口
type Achieve_interface interface {
	Ltos(c*gin.Context)
	Stol(c*gin.Context)
	Queryurl(c*gin.Context)
}
