package mydatabase

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
)
//数据库相关
var (
	Cli_redis *redis.Client
	Db *gorm.DB
)

//数据库连接相关信息
type DB_info struct{
	User string
	Password string
	Host string
	Port int
	DB_name string
}
//数据库表的设计
type Ltostable struct{
	//gorm.Model
	URL_id int64 `gorm:"primary_key"`
	//CreatedAt time.Time
	Longurl string `gorm:"type:varchar(9216)"`
	Shorturl string `gorm:"unique_index;type:varchar(32)"`
}
