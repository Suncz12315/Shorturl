package mydatabase
import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//数据库连接函数
func DB_CONN(MyUser, Password, Host,Db string ,Port int) (*gorm.DB,error){
	connArgs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", MyUser,Password, Host, Port, Db )
	db,err := gorm.Open("mysql", connArgs)
	if err != nil {
		panic("数据库连接失败")
	}
	//db.SingularTable(true)

	return db,err
}
//数据库初始化
func Init_db(){
	cn := DB_info{"root",
		"123456",
		"localhost",
		3306,
		"lkdb"}
	Db,_ = DB_CONN(cn.User,cn.Password,cn.Host,cn.DB_name,cn.Port)
	if Db.HasTable("ltostables")==false {
		Db.Set("gorm:table_options", "collate=utf8_bin").CreateTable(&Ltostable{})
	}
}
//关闭数据库
func Close_db(){
	Db.Close()
}

//redis初始化
func Init_redis(){
	Cli_redis = redis.NewClient(&redis.Options{
		Addr:      "127.0.0.1:6379",
		Password:   "",
		DB:          1,
	})
	value,_ := Cli_redis.Get("count_url").Result() //先判定计数器是否已存在
	if value == "" {
		err := Cli_redis.Set("count_url", "0", 0).Err()
		if err != nil {
			panic(err)
		}
	}
	/*pong,err :=cli_redis.Ping().Result()
	fmt.Println(pong,err)*/
}

