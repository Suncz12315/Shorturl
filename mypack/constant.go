package mypack

import "sync"

var(
	cnt int64 = 0
	stoi_maps = make(map[string]int64)//记录url->cnt(10进制)
	itos_maps = make(map[int64]string)//cnt->url 恢复长网址
	hash_maps = make(map[byte]int)
	lk sync.Mutex //创建互斥锁
	dicts string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

type InsertForm struct{
	Longurl     string `form:"longurl" json:"longurl" binding:"required"`
}

type QueryForm struct{
	Longurl     string `form:"longurl" json:"longurl" `
	Shorturl    string `form:"shorturl" json:"longurl" `
}