package mypack

import (
	"github.com/bmizerany/assert"
	"mypack/mydatabase"
	"testing"
)
func TestDB_CONN(t *testing.T) {
	_,err_db:= mydatabase.DB_CONN("root","123","localhost","lkdb",3306)
	if err_db!=nil{
		t.Log("correct")
	}
}
func TestIsvalid_url(t *testing.T) {
	test_res := Isvalid_url("www.baidu.com")
	res := false
	assert.Equal(t,test_res,res)
	test_res = Isvalid_url("http://www.baidu.com")
	res = true
	assert.Equal(t,test_res,res)
}

func TestInit_maps(t*testing.T){
	Init_maps()
	//fmt.Println(dicts[0])
	var x byte = '0'
	assert.Equal(t,dicts[0],x)
	x = 'A'
	assert.Equal(t,dicts[10],x)
	x = 'z'
	assert.Equal(t,dicts[61],x)
}
//测试10进制转62进制
func TestTo62mode(t *testing.T) {
	str:=To62mode(1)
	var test_str string = "1"
	assert.Equal(t,str,test_str)
	str = To62mode(100)
	test_str="1c"
	assert.Equal(t,str,test_str)
	str = To62mode(1000)
	test_str="G8"
	assert.Equal(t,str,test_str)
}
//测试62进制转10进制
func TestTo10mode(t *testing.T) {
    var item int64 =To10mode("www.ww")
    var test_item int64 = 0
    assert.Equal(t,item,test_item)
    item = To10mode("1")
    test_item = 1
    assert.Equal(t,item,test_item)
    item = To10mode("G8")
    test_item = 1000
    assert.Equal(t,item,test_item)
}
