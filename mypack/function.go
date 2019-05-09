package mypack
import (
	"math"
	"strings"
)
func Init_maps() map[byte]int {
	for i := 0;i<62;i++{
		hash_maps[dicts[i]] = i
	}
	return hash_maps
}
//10进制转62进制
func To62mode(num int64) string{
	var base int64 = 62
	str := ""
	for num/base!=0{
		str = string(dicts[num%base])+str
		num = num/base
	}
	str = string(dicts[num%base])+str
	return str
}
//62进制转10进制
func To10mode(s string) int64{
	size :=len(s)
	var num float64 = 0.0
	for i:=size-1;i>=0;i--{
		_,ok:=hash_maps[s[i]]//判断该短网址是否符合规则
		if(ok) {
			num += float64(hash_maps[s[i]]) * math.Pow(float64(62), float64(size-1-i))
		}else{
			return 0
		}
	}
	var x int64 = int64(num)
	return x
}

func Isvalid_url(url string) bool{
	return strings.HasPrefix(url,"http://") || strings.HasPrefix(url,"https://")
}

